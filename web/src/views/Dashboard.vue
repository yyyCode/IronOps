<template>
  <div class="dashboard">
    <div class="page-header">
      <h2>IronOps 智能运维平台</h2>
      <p class="subtitle">实时监控 · 智能预测 · 自动化运维 <span class="status-tag">● 实时数据连接中...</span></p>
    </div>

    <el-row :gutter="20" class="stat-cards">
      <el-col :span="6">
        <el-card shadow="hover">
          <div class="stat-item">
            <div class="stat-title">CPU使用率 <el-tag :type="stats.cpu_usage < 80 ? 'success' : 'danger'" size="small" effect="plain">{{ stats.cpu_usage < 80 ? '+良好' : '-高负荷' }}</el-tag></div>
            <div class="stat-value">{{ stats.cpu_usage ? stats.cpu_usage.toFixed(1) : 0 }}%</div>
            <div class="stat-chart-mini">
              <!-- Placeholder for mini chart -->
               <div class="mini-chart-line" :class="stats.cpu_usage < 80 ? 'success' : 'danger'"></div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
           <div class="stat-item">
            <div class="stat-title">系统健康度 <el-tag :type="stats.health_score > 80 ? 'success' : 'warning'" size="small" effect="plain">{{ stats.health_score > 80 ? '+健康' : '-需关注' }}</el-tag></div>
            <div class="stat-value">{{ stats.health_score ? stats.health_score.toFixed(1) : 0 }}%</div>
             <div class="stat-chart-mini">
               <div class="mini-chart-line" :class="stats.health_score > 80 ? 'success' : 'warning'"></div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
           <div class="stat-item">
            <div class="stat-title">系统响应时间 <el-tag type="success" size="small" effect="plain">-优秀</el-tag></div>
            <div class="stat-value">{{ stats.response_time }} ms</div>
             <div class="stat-chart-mini">
               <div class="mini-chart-line success"></div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
           <div class="stat-item">
            <div class="stat-title">系统稳定性 <el-tag type="success" size="small" effect="plain">-良好</el-tag></div>
            <div class="stat-value">{{ stats.stability ? stats.stability.toFixed(1) : 0 }}%</div>
             <div class="stat-chart-mini">
               <div class="mini-chart-line success"></div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" class="main-panels">
      <el-col :span="12">
        <el-card class="box-card system-info" header="系统信息">
          <el-descriptions :column="2" border>
            <el-descriptions-item label="主机名">{{ stats.host_name || '-' }}</el-descriptions-item>
            <el-descriptions-item label="操作系统">{{ stats.os || '-' }} {{ stats.platform_ver || '' }}</el-descriptions-item>
            <el-descriptions-item label="系统架构">{{ stats.platform || '-' }}</el-descriptions-item>
            <el-descriptions-item label="CPU型号">{{ stats.cpu_model || '-' }}</el-descriptions-item>
            <el-descriptions-item label="CPU核心数">{{ stats.cpu_cores || '-' }} 核</el-descriptions-item>
            <el-descriptions-item label="内存使用">{{ stats.memory_usage || 0 }}%</el-descriptions-item>
          </el-descriptions>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card class="box-card network-traffic" header="内存使用详情">
            <div ref="trafficChart" style="height: 250px;"></div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" class="bottom-panels">
      <el-col :span="24">
         <el-card header="性能趋势分析">
             <div style="height: 200px; display: flex; align-items: center; justify-content: center; color: #909399;">
                 ECharts Trend Chart Placeholder (Connect to historical data API)
             </div>
         </el-card>
      </el-col>
    </el-row>

  </div>
</template>

<script setup>
import { onMounted, onUnmounted, ref, watch } from 'vue';
import * as echarts from 'echarts';

const trafficChart = ref(null);
const stats = ref({
    cpu_usage: 0,
    health_score: 100,
    response_time: 0,
    stability: 100,
    host_name: '',
    os: '',
    platform: '',
    platform_ver: '',
    cpu_model: '',
    cpu_cores: 0,
    memory_usage: 0,
    memory_total: 0,
    memory_used: 0
});

let ws = null;
let myChart = null;

const initWebSocket = () => {
    // Determine WS protocol (ws or wss) based on current page protocol
    const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
    const host = 'localhost:8080'; // Hardcoded for dev, normally use window.location.host or env var
    
    // Get username from localStorage as simple token
    const user = JSON.parse(localStorage.getItem('user'));
    const token = user ? user.username : '';

    const wsUrl = `${protocol}//${host}/api/v1/ws/dashboard?token=${token}`;

    ws = new WebSocket(wsUrl);

    ws.onopen = () => {
        console.log("WebSocket connected");
        document.querySelector('.status-tag').textContent = "● 实时数据已连接";
        document.querySelector('.status-tag').style.color = "#67c23a";
    };

    ws.onmessage = (event) => {
        try {
            const data = JSON.parse(event.data);
            if (data) {
                stats.value = data;
                updateCharts();
            }
        } catch (e) {
            console.error("Failed to parse WS message", e);
        }
    };

    ws.onclose = () => {
        console.log("WebSocket disconnected");
        document.querySelector('.status-tag').textContent = "● 连接断开";
        document.querySelector('.status-tag').style.color = "#f56c6c";
        // Try to reconnect after 5 seconds
        setTimeout(initWebSocket, 5000);
    };

    ws.onerror = (error) => {
        console.error("WebSocket error", error);
    };
};

const updateCharts = () => {
    if (myChart && stats.value) {
        // Update memory chart
        const memoryUsedGB = (stats.value.memory_used / 1024 / 1024 / 1024).toFixed(2);
        const memoryFreeGB = ((stats.value.memory_total - stats.value.memory_used) / 1024 / 1024 / 1024).toFixed(2);
        
        myChart.setOption({
            series: [{
                data: [
                    { value: stats.value.memory_used, name: `已用: ${memoryUsedGB} GB` },
                    { value: stats.value.memory_total - stats.value.memory_used, name: `可用: ${memoryFreeGB} GB` },
                ]
            }]
        });
    }
};

onMounted(() => {
  initWebSocket();

  if (trafficChart.value) {
    myChart = echarts.init(trafficChart.value);
    const option = {
      tooltip: {
        trigger: 'item'
      },
      legend: {
        top: '5%',
        left: 'center'
      },
      series: [
        {
          name: '内存使用',
          type: 'pie',
          radius: ['40%', '70%'],
          avoidLabelOverlap: false,
          itemStyle: {
            borderRadius: 10,
            borderColor: '#fff',
            borderWidth: 2
          },
          label: {
            show: false,
            position: 'center'
          },
          emphasis: {
            label: {
              show: true,
              fontSize: 20,
              fontWeight: 'bold'
            }
          },
          labelLine: {
            show: false
          },
          data: [
            { value: 0, name: '已用' },
            { value: 0, name: '可用' },
          ]
        }
      ]
    };
    myChart.setOption(option);
  }
});

onUnmounted(() => {
    if (ws) {
        ws.close();
    }
    if (myChart) {
        myChart.dispose();
    }
});
</script>

<style scoped>
.dashboard {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.page-header h2 {
  margin: 0;
  font-size: 24px;
  color: #303133;
}

.subtitle {
  color: #909399;
  font-size: 14px;
  margin-top: 5px;
}

.status-tag {
  color: #f56c6c;
  float: right;
}

.stat-item {
  display: flex;
  flex-direction: column;
}

.stat-title {
  color: #909399;
  font-size: 14px;
  display: flex;
  justify-content: space-between;
}

.stat-value {
  font-size: 28px;
  font-weight: bold;
  margin: 10px 0;
  color: #303133;
}

.mini-chart-line {
  height: 4px;
  width: 100%;
  border-radius: 2px;
  background-color: #e4e7ed;
  position: relative;
}

.mini-chart-line.success::after {
  content: '';
  position: absolute;
  left: 0;
  top: 0;
  height: 100%;
  width: 70%;
  background-color: #67c23a;
  border-radius: 2px;
}

.mini-chart-line.warning::after {
  content: '';
  position: absolute;
  left: 0;
  top: 0;
  height: 100%;
  width: 40%;
  background-color: #e6a23c;
  border-radius: 2px;
}

.main-panels {
  margin-top: 10px;
}
.bottom-panels {
    margin-top: 10px;
}
</style>
