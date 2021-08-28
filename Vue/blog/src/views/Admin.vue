<template>
  <div id="main">
    <AdminHeader></AdminHeader>
  </div>
</template>

<script>
import AdminHeader from "../components/admin/AdminHeader";
import {loginCheck} from "../api/login";
export default {
  name: "Admin",
  components: {AdminHeader},
  async created(){
    const response = await loginCheck()
    if(response.status !==0 || response.data.privilege !== 1){
      alert("请登录管理员账号")
      window.location = '/home'
    }
  }
}
</script>

<style scoped>
div#main{
  min-width: 1440px;
}
</style>
