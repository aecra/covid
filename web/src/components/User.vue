<template>
  <el-form :model="form" label-width="120px">
    <el-form-item label="姓名">
      <el-input disabled v-model="form.username" />
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
      <el-button type="primary" @click="report">立即上报</el-button>
      <el-button @click="getUser">刷新</el-button>
    </el-form-item>
  </el-form>
</template>

<script lang="ts" setup>
import { reactive } from 'vue';
import { ElLoading } from 'element-plus';
import { ElMessage } from 'element-plus';
import DataService from '../utils/DataService';

// do not use same name with ref
let form = reactive({
  username: '',
  email: '',
  state: false,
  position: '',
  eaisess: '',
  uukey: '',
  home: '',
});

const onSubmit = async () => {
  const loadingInstance = ElLoading.service({ target: '.el-main' });
  const [err] = await DataService.updateUser(form);
  if (err) {
    ElMessage.error('更新数据失败.');
  } else {
    ElMessage.success('更新数据成功.');
  }
  loadingInstance.close();
};

const report = async () => {
  const loadingInstance = ElLoading.service({ target: '.el-main' });
  const [err] = await DataService.report();
  if (err) {
    ElMessage.error('上报失败.');
  } else {
    ElMessage.success('上报成功.');
  }
  loadingInstance.close();
}

const getUser = async () => {
  const loadingInstance = ElLoading.service({ target: '.el-main' });
  const [err, res] = await DataService.getUser();
  if (err) {
    ElMessage.error('获取数据失败.');
  }
  form = Object.assign(form, res);
  loadingInstance.close();
};

getUser();
</script>

<script lang="ts">
export default {
  name: 'UserPage',
};
</script>
