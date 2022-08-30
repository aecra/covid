<template>
  <div class="login">
    <div>
      <el-input placeholder="请输入用户名" v-model="user.username" clearable class="input_style"></el-input>
      <span v-if="error.username" class="err-msg">{{ error.username }}</span>
    </div>
    <div>
      <el-input placeholder="请输入密码" v-model="user.password" show-password class="input_style"></el-input>
      <span v-if="error.password" class="err-msg">{{ error.password }}</span>
    </div>
    <div>
      <el-button type="primary" @click="login" class="login_style">登录</el-button>
    </div>
    <div style="width: 200px">
      <el-link :underline="false" href="/#/register" style="float: right" type="primary">立即注册</el-link>
    </div>
  </div>
</template>

<script lang="ts" setup>
import auth from "../utils/auth";
import {ElLoading, ElMessage} from 'element-plus'
import {reactive} from "vue";

let user = reactive({
  "username": "",
  "password": "",
})
let error = reactive({
  "username": "",
  "password": "",
})

const USERNAME_PATTERN = /^[a-zA-Z0-9_-]{5,16}$/
const PASSWORD_PATTERN = /^[a-zA-Z0-9!@#$%&*_-]{8,16}$/

let login = () => {
  const loading = ElLoading.service({
    lock: true,
    text: 'Loading',
    background: 'rgba(0, 0, 0, 0.7)',
  })
  // init errors
  error.username = ""
  error.password = ""
  if (!USERNAME_PATTERN.test(user.username)) {
    error.username = "用户名格式不正确"
    loading.close()
    return
  }
  if (!PASSWORD_PATTERN.test(user.password)) {
    error.password = "密码格式不正确"
    loading.close()
    return
  }
  auth.login(user).then((res) => {
    if (res.status === 200) {
      localStorage.setItem("token", res.data.data.token)
      setTimeout(() => {
        window.location.href = '/'
      }, 1000)
    } else {
      ElMessage.error(res.data.data.message)
    }
  }).catch((err) => {
    ElMessage.error(err.response.data.message)
  }).finally(() => {
    loading.close()
  })
}
</script>

<script lang="ts">
export default {
  name: "LoginPage",
}
</script>

<style>
.login {
  height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background-color: #F5F5FA;
}

.login div {
  display: flex;
  flex-direction: column;
}

.err-msg {
  color: red;
  font-size: 12px;
  margin: 0 0 10px 10px;
}

.input_style, .login_style {
  width: 200px;
  margin-bottom: 10px;
}
</style>

