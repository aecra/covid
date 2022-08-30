<template>
  <el-table :data="tableData.data" border style="width: 100%">
    <el-table-column prop="time" label="上报时间"/>
    <el-table-column prop="position" label="地址"/>
    <el-table-column prop="email" label="通知邮件"/>
    <el-table-column prop="report_result" label="上报结果"/>
    <el-table-column prop="notice_result" label="通知结果"/>
  </el-table>
</template>

<script lang="ts" setup>
import DataService from "../utils/DataService";
import {reactive} from "vue";

const tableData = reactive({"data":[]})
DataService.getRecords().then((data) => {
  tableData.data = data[1];
  tableData.data.forEach((element: {
    time: string;
  }) => {
    const date = new Date(element.time)
    element.time = date.toLocaleDateString()+' '+date.getHours()+':'+date.getMinutes()+':'+date.getSeconds();
  });
})
</script>

<script lang="ts">
export default {
  name: "RecordsPage"
}
</script>

<style scoped>

</style>
