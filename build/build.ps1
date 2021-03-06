param(
  [String]$BuildDate = (Get-Date -Format "yyyy.MM.dd"),
  [String]$GitRev = "$(git rev-parse --short HEAD)",
  [String]$BuildNumber = "$([int]$env:BUILD_NUMBER)",
  [String]$ClientId,
  [String]$ClientSecret
)

$RootDir = Join-Path $PsScriptRoot ".."

Task PackageRestore {
  try {
    pushd $RootDir
    exec { go get -t -d ./... }
  } finally {
    popd
  }
}

Task GenerateMocks {
  try {
    pushd $RootDir
    exec { go get github.com/vektra/mockery/cmd/mockery }
    exec { go get golang.org/x/tools/cmd/stringer }
    exec { go generate ./... }
  } finally {
    popd
  }
}

Task Build PackageRestore, {
  try {
    pushd $RootDir

    $version="${BuildDate}-${GitRev}-${BuildNumber}"

    exec { go install -ldflags "-X github.com/waives/surf/ch360.Version=$version" ./cmd/surf }
  } finally {
    popd
  }
}

function Register-CustomAssertions([string]$path) {
  Write-Host -ForegroundColor Green "Registering custom assertions from '$path'"
  Get-ChildItem (Resolve-Path "$PsScriptRoot/../$path") -Exclude "*.Tests.ps1" |% {
    $name = ($_.Name)
    try {
      . $_
      Write-Host -ForegroundColor DarkGreen "  Registered $name"
    } catch {
      Write-Host -ForegroundColor Red "  Failed to register $name"
    }
  }
}

Task Test Build, {
  try {
    pushd $RootDir

    $vet = Allow-Vet
    exec { go test -v $vet -race ./... }

    Register-CustomAssertions "test/assertions/pester"

    $testResults = Invoke-Pester -PassThru -Script @{
      Path="test";
      Parameters = @{ClientId = $ClientId; ClientSecret = $ClientSecret}
    }
    assert ($testResults.FailedCount -eq 0)
  } finally {
    popd
  }
}

function Allow-Vet {
  $os = $PSVersionTable.OS
  if (($os -eq $null) -or ($os.StartsWith("Microsoft Windows"))) {
    Write-Warning "OS is Windows, disabling go vet (see golang/go#27089)"
    return "-vet=off"
  }

  return ""
}

Task . PackageRestore, Build
