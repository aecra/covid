<template>
  <div class="menu">
    <router-link to="/">
      <el-button class="plain" plain>统计展示</el-button>
    </router-link>
    <router-link to="/User">
      <el-button class="plain" plain>用户设置</el-button>
    </router-link>
    <router-link to="/Help">
      <el-button class="plain" plain>帮助说明</el-button>
    </router-link>
    <router-link to="/Clause">
      <el-button class="plain" plain>使用条款</el-button>
    </router-link>
  </div>
</template>


<script lang="ts" setup>

</script>

<style>
.menu {
  margin: 20px 0 0 0;
  width: 100%;
  display: flex;
  flex-wrap: wrap;
}

.plain {
  margin: 10px 10px;
}
</style>