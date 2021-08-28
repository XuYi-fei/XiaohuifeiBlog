<template>
  <el-empty description="Welcome" image="/img/bg1.png" image-size=1400></el-empty>
  <el-button type="primary" @click="goBack()">返回主页</el-button>
</template>

<script>
export default {
  name: "AdminIndex",
  methods:{
    goBack(){
      window.location = '/home'
    }
  }
}
</script>

<style scoped>

</style>
