<template>
  <div class="login">
    <div>
      <el-input placeholder="请输入用户名" v-model="user.username" clearable class="input_style"></el-input>
      <span v-if="error.username" class="err-msg">{{ error.username }}</span>
    </div>
    <div>
      <el-input placeholder="请输入邮箱" v-model="user.email" class="input_style"></el-input>
      <span v-if="error.email" class="err-msg">{{ error.email }}</span>
    </div>
    <div>
      <el-input placeholder="请输入密码" v-model="user.password" show-password class="input_style"></el-input>
      <span v-if="error.password" class="err-msg">{{ error.password }}</span>
    </div>
    <div>
      <el-input placeholder="请输入密码" v-model="user.confirmation_password" show-password
                class="input_style"></el-input>
      <span v-if="error.confirmation_password" class="err-msg">{{ error.confirmation_password }}</span>
    </div>
    <div>
      <el-button type="primary" @click="login" class="login_style">注册</el-button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import auth from "../utils/auth";
import {ElLoading, ElMessage} from 'element-plus'
import {reactive} from "vue";

let user = reactive({
  "username": "",
  "email": "",
  "password": "",
  "confirmation_password": "",
})
let error = reactive({
  "username": "",
  "email": "",
  "password": "",
  "confirmation_password": "",
})

const USERNAME_PATTERN = /^[a-zA-Z0-9_-]{5,16}$/
const EMAIL_PATTERN = /^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$/
const PASSWORD_PATTERN = /^[a-zA-Z0-9!@#$%&*_-]{8,16}$/

let login = () => {
  const loading = ElLoading.service({
    lock: true,
    text: 'Loading',
    background: 'rgba(0, 0, 0, 0.7)',
  })
  // init errors
  error.username = ""
  error.email = ""
  error.password = ""
  error.confirmation_password = ""
  if (!USERNAME_PATTERN.test(user.username)) {
    error.username = "用户名格式不正确"
    loading.close()
    return
  }
  if (!EMAIL_PATTERN.test(user.email)) {
    error.email = "邮箱格式不正确"
    loading.close()
    return
  }
  if (!PASSWORD_PATTERN.test(user.password)) {
    error.password = "密码格式不正确"
    loading.close()
    return
  }
  if (user.password !== user.confirmation_password) {
    error.confirmation_password = "两次密码不一致"
    loading.close()
    return
  }
  auth.register(user).then((res) => {
    if (res.status === 200) {
      ElMessage.success('注册成功')
      setTimeout(() => {
        window.location.href = '/#/login'
      }, 1000)
    } else {
      ElMessage.error('注册失败')
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
  name: "RegisterPage",
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

