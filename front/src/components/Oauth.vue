<script lang="ts" setup>
import { useRoute } from 'vue-router'
import { Oauth } from '../tools/oauth'
import { ElLoading } from 'element-plus'

const route = useRoute()

const loadingInstance = ElLoading.service({ fullscreen: true })
loadingInstance.setText('Verifying access token...')
Oauth(route.query)
</script>

<template></template>

<style>
</style>