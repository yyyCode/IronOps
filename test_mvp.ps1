# 1. Register Admin
try {
    $admin = Invoke-RestMethod -Method Post -Uri "http://localhost:8080/api/v1/register" -Body (@{username="admin"; password="123"; role="admin"} | ConvertTo-Json) -ContentType "application/json" -ErrorAction Stop
    Write-Host "Registered Admin: $($admin.username)"
} catch {
    Write-Host "Registration failed or user exists: $_"
}

# 2. Create Service (Header X-User: admin)
try {
    $service = Invoke-RestMethod -Method Post -Uri "http://localhost:8080/api/v1/services" -Headers @{"X-User"="admin"} -Body (@{name="OrderService"; owner="TeamA"; env="prod"} | ConvertTo-Json) -ContentType "application/json"
    Write-Host "Created Service: $($service.name) ID: $($service.ID)"

    # 3. Add Instance
    $instance = Invoke-RestMethod -Method Post -Uri "http://localhost:8080/api/v1/instances" -Headers @{"X-User"="admin"} -Body (@{service_id=$service.ID; ip="192.168.1.101"; port=8080; status="running"} | ConvertTo-Json) -ContentType "application/json"
    Write-Host "Added Instance: $($instance.ip) ID: $($instance.ID)"

    # 4. Report Metric (CPU High)
    Invoke-RestMethod -Method Post -Uri "http://localhost:8080/api/v1/metrics" -Headers @{"X-User"="admin"} -Body (@{instance_id=$instance.ID; cpu=95.5; memory=60.0} | ConvertTo-Json) -ContentType "application/json"
    Write-Host "Reported High CPU Metric"

    # 5. List Alerts
    $alerts = Invoke-RestMethod -Method Get -Uri "http://localhost:8080/api/v1/alerts" -Headers @{"X-User"="admin"}
    Write-Host "Alerts Found: $($alerts.Count)"
    $alerts | ForEach-Object { Write-Host " - Alert: $($_.message)" }

    # 6. Control Instance (Stop)
    Invoke-RestMethod -Method Post -Uri "http://localhost:8080/api/v1/instances/$($instance.ID)/control" -Headers @{"X-User"="admin"} -Body (@{action="stop"} | ConvertTo-Json) -ContentType "application/json"
    Write-Host "Stopped Instance"

    # 7. Check Audit Log
    $logs = Invoke-RestMethod -Method Get -Uri "http://localhost:8080/api/v1/audits" -Headers @{"X-User"="admin"}
    Write-Host "Audit Logs: $($logs.Count)"
    $logs | Select-Object -First 5 | ForEach-Object { Write-Host " - [$($_.action)] $($_.target) by $($_.user)" }

} catch {
    Write-Host "Error: $_"
}
