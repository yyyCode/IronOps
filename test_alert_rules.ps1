# 1. Login as Admin/Ops to get token (or just use X-User header for now since we have RoleMiddleware but maybe no strict JWT yet? Ah, AuthMiddleware checks token. Let's see if we can bypass or need to login)
# The current middleware checks for "Authorization" header or query param.
# Wait, my previous test_mvp.ps1 used `Headers @{"X-User"="admin"}`?
# Let's check `middleware/auth.go`. If it's not strictly enforcing JWT, X-User might work if I didn't change it.
# Actually, I haven't seen `middleware/auth.go`. Let's assume I need to login if I want to be safe, OR I can just try the X-User trick if the middleware supports it (Dev mode).
# Let's check `middleware/auth.go` first to be sure.

# Just in case, I will try to Register/Login first.

$ErrorActionPreference = "Stop"

function Invoke-Api {
    param($Method, $Uri, $Body, $Headers)
    try {
        $params = @{
            Method = $Method
            Uri = $Uri
            ContentType = "application/json"
        }
        if ($Body) { $params.Body = $Body }
        if ($Headers) { $params.Headers = $Headers }
        return Invoke-RestMethod @params
    } catch {
        Write-Host "Error calling $Uri : $($_.Exception.Message)"
        if ($_.Exception.Response) {
            $reader = New-Object System.IO.StreamReader($_.Exception.Response.GetResponseStream())
            Write-Host "Response: $($reader.ReadToEnd())"
        }
        throw $_
    }
}

# 1. Login (Mock Login returns user info, assuming it sets a cookie or token? The code I saw earlier just returned JSON. 
# Middleware usually expects a header.
# Let's check `middleware/auth.go` quickly.
# ... I'll just assume X-User works for now based on previous context or I'll use the result of Login if it returns a token.
# Wait, the `LoginHandler` I saw: `c.JSON(http.StatusOK, gin.H{"message": "login success", "user": user})`. It didn't seem to return a token.
# Maybe AuthMiddleware is dummy?

# Let's try creating a rule with X-User: admin.

Write-Host "Creating Alert Rule..."
$ruleBody = @{
    name = "High CPU Test"
    metric_type = "cpu"
    condition = ">"
    threshold = 10.0
    severity = "critical"
    enabled = $true
    description = "CPU is too high"
} | ConvertTo-Json

$headers = @{"X-User"="admin"; "X-Role"="admin"} # Assuming middleware trusts this for dev

# Try to bypass auth or assume it works
try {
    # Creating Rule
    $rule = Invoke-Api -Method Post -Uri "http://localhost:8080/api/v1/alert-rules" -Body $ruleBody -Headers $headers
    Write-Host "Rule Created: ID=$($rule.ID)"
} catch {
    Write-Host "Failed to create rule. Maybe Auth issue."
    exit 1
}

# 2. Report Metric that violates rule
# Need an instance ID. Let's list services first to find one or create one.
Write-Host "Listing Services..."
$services = Invoke-Api -Method Get -Uri "http://localhost:8080/api/v1/services" -Headers $headers
if ($services.Count -eq 0) {
    # Create Service & Instance
    Write-Host "Creating Service..."
    $svc = Invoke-Api -Method Post -Uri "http://localhost:8080/api/v1/services" -Body (@{name="TestSvc"; owner="Me"; env="dev"} | ConvertTo-Json) -Headers $headers
    Write-Host "Creating Instance..."
    $inst = Invoke-Api -Method Post -Uri "http://localhost:8080/api/v1/instances" -Body (@{service_id=$svc.ID; ip="1.2.3.4"; port=80; status="running"} | ConvertTo-Json) -Headers $headers
    $instanceID = $inst.ID
} else {
    $instanceID = $services[0].Instances[0].ID
}

Write-Host "Reporting Metric for Instance $instanceID (CPU=15.0)..."
Invoke-Api -Method Post -Uri "http://localhost:8080/api/v1/metrics" -Body (@{instance_id=$instanceID; cpu=15.0; memory=50.0} | ConvertTo-Json) -Headers $headers

# 3. Check Alerts
Start-Sleep -Seconds 1
Write-Host "Checking Alerts..."
$alerts = Invoke-Api -Method Get -Uri "http://localhost:8080/api/v1/alerts" -Headers $headers

$found = $false
foreach ($a in $alerts) {
    if ($a.message -like "*CPU is too high*") {
        Write-Host "SUCCESS: Found alert -> $($a.message)"
        $found = $true
        break
    }
}

if (-not $found) {
    Write-Host "FAILURE: Did not find expected alert."
}
