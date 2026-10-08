param(
  [string]$RawDir = "raw",
  [string]$ExtractedDir = "extracted",
  [int]$ShardIndex = 0,
  [int]$ShardCount = 1
)
$ErrorActionPreference = "Stop"
if($ShardCount -lt 1 -or $ShardIndex -lt 0 -or $ShardIndex -ge $ShardCount){
  throw "ShardIndex must be in [0, ShardCount)."
}
$RawDir = [IO.Path]::GetFullPath($RawDir)
$ExtractedDir = [IO.Path]::GetFullPath($ExtractedDir)
New-Item -ItemType Directory -Force -Path $ExtractedDir | Out-Null
$word = New-Object -ComObject Word.Application
$excel = New-Object -ComObject Excel.Application
$word.Visible = $false
$word.DisplayAlerts = 0
$word.AutomationSecurity = 3
$excel.Visible = $false
$excel.DisplayAlerts = $false
$excel.AutomationSecurity = 3
try {
  $files = @(Get-ChildItem -LiteralPath $RawDir -File | Sort-Object Name)
  for($fileIndex = $ShardIndex; $fileIndex -lt $files.Count; $fileIndex += $ShardCount) {
    $_ = $files[$fileIndex]
    $sourcePath = $_.FullName
    $target = [IO.Path]::GetFullPath((Join-Path $ExtractedDir ($_.BaseName + ".txt")))
    if(Test-Path -LiteralPath $target){ continue }
    $doc = $null
    $book = $null
    try {
      if(Test-Path -LiteralPath ($target+".error")){ Remove-Item -LiteralPath ($target+".error") -Force }
      switch ($_.Extension.ToLowerInvariant()) {
        ".doc" { $doc=$word.Documents.Open([string]$sourcePath,$false,$true,$false,"","",$true);$doc.SaveAs2([string]$target,2) }
        ".docx" { $doc=$word.Documents.Open([string]$sourcePath,$false,$true,$false,"","",$true);$doc.SaveAs2([string]$target,2) }
        ".xlsx" {
          $book=$excel.Workbooks.Open([string]$sourcePath,0,$true)
          $lines=New-Object System.Collections.Generic.List[string]
          foreach($sheet in $book.Worksheets){
            $lines.Add("# sheet: " + [string]$sheet.Name)
            $used=$sheet.UsedRange
            $values=$used.Value2
            if($values -is [Array]){
              for($r=1;$r -le $values.GetLength(0);$r++){
                $cells=for($c=1;$c -le $values.GetLength(1);$c++){[string]$values[$r,$c]}
                $lines.Add(($cells -join "`t"))
              }
            } elseif($null -ne $values) {
              $lines.Add([string]$values)
            }
            [void][Runtime.InteropServices.Marshal]::FinalReleaseComObject($used)
            [void][Runtime.InteropServices.Marshal]::FinalReleaseComObject($sheet)
          }
          [IO.File]::WriteAllLines($target,$lines,[Text.UTF8Encoding]::new($false))
        }
      }
    } catch {
      [IO.File]::WriteAllText($target+".error",$_.Exception.ToString())
    } finally {
      if($null -ne $doc){ $doc.Close($false); [void][Runtime.InteropServices.Marshal]::FinalReleaseComObject($doc) }
      if($null -ne $book){ $book.Close($false); [void][Runtime.InteropServices.Marshal]::FinalReleaseComObject($book) }
    }
  }
} finally {
  $word.Quit()
  $excel.Quit()
  [void][Runtime.InteropServices.Marshal]::FinalReleaseComObject($word)
  [void][Runtime.InteropServices.Marshal]::FinalReleaseComObject($excel)
  [GC]::Collect()
  [GC]::WaitForPendingFinalizers()
}
