<template>
  <div class="collapse">
    <el-collapse v-model="activeNames">
      <el-collapse-item title="邮箱" name="1">
        <div>邮箱用来发送疫情打卡结果通知。</div>
      </el-collapse-item>
      <el-collapse-item title="位置" name="2">
        <div>位置有两个选项，一个是 School，一个是 Home。</div>
        <div>
          如果你选择了School，那么将进行晨午检打卡。
          如果选择了Home，那么将根据“Home 数据”进行健康卡打卡。
        </div>
      </el-collapse-item>
      <el-collapse-item title="状态" name="3">
        <div>可以选择是否开启疫情打卡。</div>
      </el-collapse-item>
      <el-collapse-item title="Cookie eaisess" name="4">
        <div>该 Cookie 字段用于记录登录状态。在这里用于打卡的认证。</div>
        <div>
          获取方法以Edge浏览器为例，在微信中打开疫情打卡的网址，
          选择在浏览器中打开，此时会跳转到学校的登录页面，输入您的账号密码登录后，
          按 F12 打开开发者工具， 在应用程序 tab 栏中找到存储，展开本地存储，
          选中学校网址，此时能看到 Cookie 字段。
        </div>
      </el-collapse-item>
      <el-collapse-item title="Cookie uukey" name="5">
        <div>参见 Cookie eaisess。</div>
      </el-collapse-item>
      <el-collapse-item title="Home 数据" name="6">
        <div>这个字段为在家中的打卡数据，如果你选择了Home，那么将根据这个字段进行健康卡打卡。</div>
        <div>该字段获取方法参见 Cookie eaisess。</div>
        <div>在打开开发者工具后，打开网络 Tab，填写页面中疫情打卡信息，然后选择提交，此时最后的网络请求即为该字段所需内容</div>
      </el-collapse-item>
    </el-collapse>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { verify } from '../tools/oauth'

verify()
const activeNames = ref(['1'])

</script>


<style>
.collapse {
  max-width: 960px;
}
</style>