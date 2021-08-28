<template>
  <suspense>
    <template #default>
      <AdminBlogsPage />
    </template>
    <template #fallback>
      加载中...
    </template>
  </suspense>

</template>

<script>
import {getAllBlogsTmp} from "../../api/admin";
import AdminBlogsPage from "./AdminBlogsPage";

export default {
  name: "AdminBlogs",
  components: {AdminBlogsPage},
  data(){
    return{
      blogs: [],
      empty: false,
    }
  },
}
</script>

<style scoped>

</style>
