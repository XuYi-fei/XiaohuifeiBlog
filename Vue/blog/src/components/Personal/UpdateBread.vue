<template>
<span class="el-breadcrumb__item" data-v-b8276df4="">
  <span class="el-breadcrumb__inner is-link" role="link" @click="back()">个人空间</span>
  <i class="el-breadcrumb__separator el-icon-arrow-right"></i>
</span>
<span class="el-breadcrumb__item" data-v-b8276df4="">
  <span class="el-breadcrumb__inner is-link" role="link">更新博客</span>
  <i class="el-breadcrumb__separator el-icon-arrow-right"></i>
</span>
</template>

<script>
export default {
  name: "UpdateBread",
  methods:{
    back(){
      history.go(-1)
    }
  }
}
</script>

<style scoped>

</style>
