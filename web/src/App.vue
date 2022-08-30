<template>
  <div class="common-layout">
    <el-container class="full-height" v-if="$route.meta.keepalive">
      <el-header class="header">
        <keep-alive>
          <v_Header></v_Header>
        </keep-alive>
      </el-header>
      <el-container>
        <el-aside class="hidden-sm-and-down" width="200px">
          <keep-alive>
            <v_Menu class="full-height"></v_Menu>
          </keep-alive>
        </el-aside>
        <el-main class="main">
          <MenuInside class="hidden-md-and-up"></MenuInside>
          <el-scrollbar class="in-main" max-height="">
            <keep-alive>
              <router-view></router-view>
            </keep-alive>
          </el-scrollbar>
        </el-main>
      </el-container>
    </el-container>
    <router-view v-if="!$route.meta.keepalive"></router-view>
  </div>
</template>

<script lang="ts">
import Menu from './components/Menu.vue'
import MenuInside from './components/MenuInside.vue'
import Header from "./components/Header.vue"
export default {
  name: 'CommonComponent',
  components: {
    v_Header: Header,
    MenuInside: MenuInside,
    v_Menu: Menu,
  }
}
</script>

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
  padding: 20px 20px 80px;
}
</style>

