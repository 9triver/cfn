<script setup lang="ts">

</script>

<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<style scoped>
:root {
  /*--graph-background: #F2F7FA;*/
  /*--text-color: var(--el-text-color-primary);
  --brand-color: var(--el-color-primary);*/
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  /* text-align: center; */
  color: #2c3e50;
  width: 100%; /* 全屏宽度 */
  height: 100vh; /* 全屏高度 */
}

body, html {
  overflow: hidden;
  margin: 0;
  padding: 0;
}

</style>
