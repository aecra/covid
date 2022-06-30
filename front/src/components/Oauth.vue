<script lang="ts" setup>
import { useRoute } from 'vue-router'
import OauthService from '../tools/oauth'
import { ElLoading } from 'element-plus'

const route = useRoute()

const loadingInstance = ElLoading.service({ fullscreen: true })
loadingInstance.setText('Verifying access token...')
OauthService.Oauth(route.query)
</script>

<template></template>

<style>
</style>