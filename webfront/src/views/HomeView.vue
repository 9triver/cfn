<script setup lang="ts">
import {onMounted, reactive, ref} from 'vue';
import {
  ApiOutlined,
  ClusterOutlined,
  ContainerOutlined,
  DashboardOutlined,
  DeploymentUnitOutlined
} from '@ant-design/icons-vue';


const selectedKeys = ref(['overview']);
const version = ref('1.22.3');
const commitId = ref('a9f5d63');

// 集群资源数据
const clusterResources = reactive({
  cpu: 32,
  memory: 128,
  nodes: 5
});

// 可用资源数据
const availableResources = reactive({
  cpu: 8,
  memory: 32,
  pods: 120
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

// 模拟从API获取数据
onMounted(() => {
  // 这里可以添加API调用逻辑
  // fetchClusterData();
});

</script>

<template>
  <a-layout class="k8s-dashboard" has-sider>
    <!-- 左侧导航栏 -->
    <a-layout-sider theme="dark" width="220"
                    :style="{ overflow: 'auto', height: '100vh', position: 'fixed', left: 0, top: 0, bottom: 0 }">
      <div class="logo">Node Dashboard</div>
      <a-menu theme="dark" mode="inline" v-model:selectedKeys="selectedKeys">
        <a-menu-item key="overview">
          <template #icon>
            <dashboard-outlined/>
          </template>
          <span>Overview</span>
        </a-menu-item>
        <a-menu-item key="nodes">
          <template #icon>
            <cluster-outlined/>
          </template>
          <span>Nodes</span>
        </a-menu-item>
        <a-menu-item key="pods">
          <template #icon>
            <container-outlined/>
          </template>
          <span>Pods</span>
        </a-menu-item>
        <a-menu-item key="deployments">
          <template #icon>
            <deployment-unit-outlined/>
          </template>
          <span>Deployments</span>
        </a-menu-item>
        <a-menu-item key="services">
          <template #icon>
            <api-outlined/>
          </template>
          <span>Services</span>
        </a-menu-item>
      </a-menu>
    </a-layout-sider>

    <!-- 右侧内容区 -->
    <a-layout :style="{ marginLeft: '220px' }">
      <a-layout-header class="header">
        <div class="version-info">
          Kubernetes Version: {{ version }}
          <span class="commit-id">Commit: {{ commitId }}</span>
        </div>
      </a-layout-header>

      <a-layout-content class="content">
        <!-- 资源概览卡片 -->
        <div class="resource-cards">
          <a-row :gutter="16">
            <a-col :span="6">
              <a-card title="Cluster Resources" class="resource-card">
                <div class="resource-item">
                  <div class="resource-label">Total CPU</div>
                  <div class="resource-value">{{ clusterResources.cpu }} cores</div>
                </div>
                <div class="resource-item">
                  <div class="resource-label">Total Memory</div>
                  <div class="resource-value">{{ clusterResources.memory }} GiB</div>
                </div>
                <div class="resource-item">
                  <div class="resource-label">Total Nodes</div>
                  <div class="resource-value">{{ clusterResources.nodes }}</div>
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
                  <div class="resource-value">{{ availableResources.memory }} GiB</div>
                </div>
                <div class="resource-item">
                  <div class="resource-label">Available Pods</div>
                  <div class="resource-value">{{ availableResources.pods }}</div>
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
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<style scoped>
.k8s-dashboard {
  min-height: 100vh;
  width: 100%;
}

.logo {
  height: 32px;
  margin: 16px;
  color: white;
  font-size: 18px;
  font-weight: bold;
  text-align: center;
}

.header {
  background: #fff;
  padding: 0 24px;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);
}

.version-info {
  font-size: 14px;
  color: #666;
}

.commit-id {
  margin-left: 12px;
  color: #999;
}

.content {
  margin: 24px;
}

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
