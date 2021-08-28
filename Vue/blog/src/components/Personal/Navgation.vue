<template>
  <el-col :span="4" style="min-width: 150px; min-height: 800px; flex-wrap: nowrap ;background-color: #545C64">
    <el-menu
        :uniqueOpened="true"
        default-active="personal-info"
        class="el-menu-vertical-demo"
        @open="handleOpen"
        @close="handleClose"
        background-color="#545c64"
        text-color="#fff"
        active-text-color="#ffd04b"
        :router="true">
      <el-menu-item index="personal-info">
        <i class="el-icon-menu"></i>
        <template #title>个人主页</template>
      </el-menu-item>
      <el-menu-item index="personal-detail">
        <i class="el-icon-menu"></i>
        <template #title>个人信息</template>
      </el-menu-item>
      <el-menu-item index="inbox">
        <i class="el-icon-box"></i>
        <template #title>收件信箱</template>
      </el-menu-item>
      <el-menu-item index="createBlog">
        <i class="el-icon-document"></i>
        <template #title>发表文章</template>
      </el-menu-item>
      <el-menu-item index="editBlog">
        <i class="el-icon-setting"></i>
        <template #title>编辑文章</template>
      </el-menu-item>
      <el-menu-item index="editSentence">
        <i class="el-icon-setting"></i>
        <template #title>编辑美句</template>
      </el-menu-item>
    </el-menu>
  </el-col>
</template>

<script>
export default {
  name: "Navgation",
  methods:{
    handleOpen(key, keyPath) {
      console.log(key, keyPath);
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath);
    }
  }
}
</script>

<style scoped>

</style>
