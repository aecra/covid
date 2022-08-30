<template>
  <el-menu>
    <router-link to="/">
      <el-menu-item index="1">
        <span class="title">统计展示</span>
      </el-menu-item>
    </router-link>
    <router-link to="/User">
      <el-menu-item index="2">
        <span class="title">用户设置</span>
      </el-menu-item>
    </router-link>
    <router-link to="/Records">
      <el-menu-item index="3">
        <span class="title">填报记录</span>
      </el-menu-item>
    </router-link>
    <router-link to="/Help">
      <el-menu-item index="4">
        <span class="title">帮助说明</span>
      </el-menu-item>
    </router-link>
    <router-link to="/Clause">
      <el-menu-item index="5">
        <span class="title">使用条款</span>
      </el-menu-item>
    </router-link>
  </el-menu>
</template>

<script lang="ts">
export default {
  name: 'MenuComponent',
}
</script>

<style>
.title {
  margin: auto;
  font-size: 16px;
  color: #666;
  line-height: 32px;
  text-decoration: none;
}
a {
  text-decoration: none;
}
</style>
