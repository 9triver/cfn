<!--suppress VueUnrecognizedSlot -->
<script setup lang="ts">
import {ref, watch} from 'vue';
import {
  ApiOutlined,
  ClusterOutlined,
  ContainerOutlined,
  DashboardOutlined,
  DeploymentUnitOutlined
} from '@ant-design/icons-vue';
import {useRoute} from "vue-router";


const selectedKeys = ref(['overview']);
const version = ref('1.22.3');
const commitId = ref('a9f5d63');

const route = useRoute();

// 监听路由变化
watch(() => route.path, (newPath) => {
  // 根据路由路径设置选中的菜单项
  selectedKeys.value = [newPath.substring(newPath.lastIndexOf("/") + 1)]
}, {immediate: true})

</script>

<template>
  <a-layout class="node-dashboard" has-sider>
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
        <router-view/>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<style scoped>
.node-dashboard {
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


</style>
