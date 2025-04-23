<!--suppress VueUnrecognizedSlot -->
<script setup lang="ts">
// 集群资源数据
import {onMounted, reactive} from "vue";
import {getAvailableResourceCount, getTotalResourceCount} from "@/api/overview";

const totalResources = reactive({
  cpu: "NaN",
  memory: "NaN",
});

// 可用资源数据
const availableResources = reactive({
  cpu: "NaN",
  memory: "NaN",
});

// 工作负载数据
const workloads = reactive({
  runningPods: 80,
  deployments: 15,
  services: 10
});

// 集群健康数据
const health = reactive({
  nodesReadyPercent: 80,
  podsRunningPercent: 95,
  cpuUsagePercent: 65
});

// 节点列表数据
const nodes = reactive([
  {
    key: '1',
    name: 'node-1',
    status: 'Ready',
    cpuCapacity: 8,
    cpuUsed: 5,
    memoryCapacity: 32,
    memoryUsed: 24,
    pods: 20
  },
  {
    key: '2',
    name: 'node-2',
    status: 'Ready',
    cpuCapacity: 8,
    cpuUsed: 6,
    memoryCapacity: 32,
    memoryUsed: 28,
    pods: 18
  },
  {
    key: '3',
    name: 'node-3',
    status: 'NotReady',
    cpuCapacity: 8,
    cpuUsed: 2,
    memoryCapacity: 32,
    memoryUsed: 8,
    pods: 5
  },
  {
    key: '4',
    name: 'node-4',
    status: 'Ready',
    cpuCapacity: 8,
    cpuUsed: 4,
    memoryCapacity: 32,
    memoryUsed: 16,
    pods: 15
  },
  {
    key: '5',
    name: 'node-5',
    status: 'Ready',
    cpuCapacity: 8,
    cpuUsed: 7,
    memoryCapacity: 32,
    memoryUsed: 30,
    pods: 22
  }
]);

// 节点表格列定义
const nodeColumns = reactive([
  {
    title: 'Node Name',
    dataIndex: 'name',
    key: 'name'
  },
  {
    title: 'Status',
    dataIndex: 'status',
    key: 'status',
    slots: {customRender: 'status'}
  },
  {
    title: 'CPU Usage',
    key: 'cpu',
    slots: {customRender: 'cpu'}
  },
  {
    title: 'Memory Usage',
    key: 'memory',
    slots: {customRender: 'memory'}
  },
  {
    title: 'Running Pods',
    dataIndex: 'pods',
    key: 'pods'
  }
]);

// 调用 API 获取数据
onMounted(() => {
  getTotalResourceCount().then((response) => {
    // TODO: 全局请求拦截
    if (response.status === 200) {
      // console.log(response);
      const totalResourceCount = response.data.result;
      totalResources.cpu = totalResourceCount.cpu.cores;
      totalResources.memory = totalResourceCount.memory.capacity;
    }
    // TODO: 处理异常
  });
  getAvailableResourceCount().then((response) => {
    if (response.status === 200) {
      const availableResourceCount = response.data.result;
      availableResources.cpu = availableResourceCount.cpu.cores;
      availableResources.memory = availableResourceCount.memory.capacity;
    }
  })
});

</script>

<template>
  <div>
    <!-- 资源概览卡片 -->
    <div class="resource-cards">
      <a-row :gutter="16">
        <a-col :span="6">
          <a-card title="Total Resources" class="resource-card">
            <div class="resource-item">
              <div class="resource-label">Total CPU</div>
              <div class="resource-value">{{ totalResources.cpu }} cores</div>
            </div>
            <div class="resource-item">
              <div class="resource-label">Total Memory</div>
              <div class="resource-value">{{ totalResources.memory }}</div>
            </div>
          </a-card>
        </a-col>

        <a-col :span="6">
          <a-card title="Available Resources" class="resource-card">
            <div class="resource-item">
              <div class="resource-label">Available CPU</div>
              <div class="resource-value">{{ availableResources.cpu }} cores</div>
            </div>
            <div class="resource-item">
              <div class="resource-label">Available Memory</div>
              <div class="resource-value">{{ availableResources.memory }}</div>
            </div>
          </a-card>
        </a-col>

        <a-col :span="6">
          <a-card title="Running Workloads" class="resource-card">
            <div class="resource-item">
              <div class="resource-label">Running Pods</div>
              <div class="resource-value">{{ workloads.runningPods }}</div>
            </div>
            <div class="resource-item">
              <div class="resource-label">Deployments</div>
              <div class="resource-value">{{ workloads.deployments }}</div>
            </div>
            <div class="resource-item">
              <div class="resource-label">Services</div>
              <div class="resource-value">{{ workloads.services }}</div>
            </div>
          </a-card>
        </a-col>

        <a-col :span="6">
          <a-card title="Cluster Health" class="resource-card">
            <div class="resource-item">
              <div class="resource-label">Nodes Ready</div>
              <a-progress
                  :percent="health.nodesReadyPercent"
                  :status="health.nodesReadyPercent >= 90 ? 'success' : 'normal'"
              />
            </div>
            <div class="resource-item">
              <div class="resource-label">Pods Running</div>
              <a-progress
                  :percent="health.podsRunningPercent"
                  :status="health.podsRunningPercent >= 90 ? 'success' : 'normal'"
              />
            </div>
            <div class="resource-item">
              <div class="resource-label">CPU Usage</div>
              <a-progress
                  :percent="health.cpuUsagePercent"
                  :status="health.cpuUsagePercent <= 70 ? 'success' : 'exception'"
              />
            </div>
          </a-card>
        </a-col>
      </a-row>
    </div>

    <!-- 节点资源列表 -->
    <a-card title="Node Resources" class="node-list">
      <a-table
          :columns="nodeColumns"
          :data-source="nodes"
          :pagination="false"
          size="middle"
      >
        <template #status="{ text }">
          <a-tag :color="text === 'Ready' ? 'green' : 'red'">
            {{ text }}
          </a-tag>
        </template>
        <template #cpu="{ record }">
          <a-progress
              :percent="(record.cpuUsed / record.cpuCapacity * 100).toFixed(1)"
              :status="(record.cpuUsed / record.cpuCapacity) <= 0.7 ? 'success' : 'exception'"
          />
        </template>
        <template #memory="{ record }">
          <a-progress
              :percent="(record.memoryUsed / record.memoryCapacity * 100).toFixed(1)"
              :status="(record.memoryUsed / record.memoryCapacity) <= 0.7 ? 'success' : 'exception'"
          />
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<style scoped>
.resource-cards {
  margin-bottom: 24px;
}

.resource-card {
  margin-bottom: 16px;
}

.resource-item {
  margin-bottom: 12px;
}

.resource-label {
  font-size: 14px;
  color: #666;
  margin-bottom: 4px;
}

.resource-value {
  font-size: 18px;
  font-weight: 500;
}

.node-list {
  width: 100%;
  margin-top: 24px;
}
</style>