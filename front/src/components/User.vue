<template>
  <el-form :model="form" label-width="120px">
    <el-form-item label="姓名">
      <el-input disabled v-model="form.name" />
    </el-form-item>
    <el-form-item label="邮箱">
      <el-input v-model="form.email" />
    </el-form-item>
    <el-form-item label="位置">
      <el-select v-model="form.position" placeholder="请选择当前位置">
        <el-option label="School" value="school" />
        <el-option label="Home" value="home" />
      </el-select>
    </el-form-item>
    <el-form-item label="状态">
      <el-switch v-model="form.state" />
    </el-form-item>
    <el-form-item label="Cookie eaisess">
      <el-input v-model="form.eaisess" />
    </el-form-item>
    <el-form-item label="Cookie uukey">
      <el-input v-model="form.uukey" />
    </el-form-item>
    <el-form-item label="Home 数据">
      <el-input v-model="form.home" type="textarea" :autosize="{ minRows: 5, maxRows: 25 }" />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">保存修改</el-button>
      <el-button @click="getUser">刷新</el-button>
    </el-form-item>
  </el-form>
</template>

<script lang="ts" setup>
import { reactive } from 'vue'
import { ElLoading } from 'element-plus'
import { ElMessage } from 'element-plus'
import request from '../tools/request'


// do not use same name with ref
const form = reactive({
  name: '',
  email: '',
  state: false,
  position: '',
  eaisess: '',
  uukey: '',
  home: '',
})

const onSubmit = async () => {
  const data = JSON.stringify({
    name: form.name,
    email: form.email,
    position: form.position,
    state: form.state ? "on" : "off",
    eaisess: form.eaisess,
    uukey: form.uukey,
    home: form.home,
  });

  const loadingInstance = ElLoading.service({target: '.el-main'})
  try {
    await request.post('/UpdateUser', data);
  } catch {
    ElMessage.error('更新数据失败.')
  }
  loadingInstance.close()
}

const getUser = async () => {
  const loadingInstance = ElLoading.service({target: '.el-main'})
  try {
    const response = await request.post('/User');
    const data = response.data;
    form.name = data.name;
    form.email = data.email;
    form.state = (data.state === "on");
    form.position = data.position;
    form.eaisess = data.eaisess;
    form.uukey = data.uukey;
    form.home = data.home;
  } catch {
    ElMessage.error('获取数据失败.')
  }
  loadingInstance.close()
}

getUser()
</script>
