<template>
  <div id="main">
    <header> 全部文章</header>
    <el-divider style="margin: 4px"></el-divider>
    <suspense>
      <template #default>
        <EditBlogPage />
      </template>
      <template #fallback>
        加载中....
      </template>
    </suspense>

  </div>
</template>

<script>
import {deleteBlog, getBlogByUser} from "../../api/editBlog";
import {loginCheck} from "../../api/login";
import EditBlogPage from "./EditBlogPage";

export default {
  name: "EditBlog",
  components: {EditBlogPage},
}
</script>

<style scoped>
a{
  cursor: pointer;
}

div#main{
  padding: 10px;
  min-width: 600px;
  flex: 80%;
  background-color: #e6faff;
  display: flex;
  text-align: left;
  flex-direction: column;
}


</style>
