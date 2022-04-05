<script setup lang="ts">
import Menu from './Menu.vue'
import MenuInside from './Menu-inside.vue'
import Header from "./Header.vue"

</script>

<template>
  <div class="common-layout">
    <el-container class="full-height">
      <el-header class="header">
        <Header></Header>
      </el-header>
      <el-container>
        <el-aside class = "hidden-sm-and-down" width="200px">
          <Menu class="full-height"></Menu>
        </el-aside>
        <el-main class="main">
          <MenuInside class = "hidden-md-and-up"></MenuInside>
          <el-scrollbar class="in-main" max-height="">
            <slot></slot>
          </el-scrollbar>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<style>
html,
body,
#app,
.full-height {
  height: 100%;
}
.common-layout {
  height: 100vh;
  overflow: hidden;
}
.header {
  border-bottom: 1px solid #dcdcdc;
}
.main {
  padding: 0;
  height: 100vh;
}
.in-main {
  box-sizing: border-box;
  padding: 20px;
  padding-bottom: 80px;
}
</style>