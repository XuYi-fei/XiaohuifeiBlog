<template>
  <el-empty description="暂无数据" style="height: 980px" v-show="empty"></el-empty>
  <ul>
    <li v-for="item in blogs">
      <div id="block">
        <div id="arrangement">
          <p class="blog-title">{{item.title}}</p>
          <br>
          <p class="blog-user">Author: {{item.author}} &nbsp;Time: {{item.createdTime.split("+")[0].split(".")[0].replace("T", " ")}}</p>
          <p class="blog-field">Field: {{ item.field }}</p>
          <p class="blog-classification">Classification: {{ item.classification }}</p>
          <p class="blog-tag">Tag: {{ item.tag }}</p>
        </div>
        <div id="operation">
            <span>
              <el-button @click="checkBlog(item)">查看详情</el-button>
              <el-popconfirm title="确定删除吗？" @confirm="remove(item)">
                <template #reference>
                <el-button>删除</el-button>
                </template>
              </el-popconfirm>
              <el-popconfirm title="确定发布吗？" @confirm="pass(item)">
                <template #reference>
                <el-button>发布</el-button>
                </template>
              </el-popconfirm>
            </span>
        </div>
      </div>
      <el-divider style="margin: 10px; width: inherit"></el-divider>
    </li>
  </ul>
  <el-pagination
      @size-change="handleSizeChange"
      @current-change="getPageBlogs"
      v-model:currentPage="currentPage"
      layout="prev, pager, next, jumper"
      :page-count="pageNum"
      :key="pageNum"
  >
  </el-pagination>
</template>

<script>
import {defineComponent, reactive, ref} from "vue";
import {DeleteBlogs, PassBlogs} from "../../api/admin";
import {deleteBlog, getBlogPageNum, getLatestBlogs} from "../../api/editBlog";

export default defineComponent({
  name: "AdminBlogsAllPage",
  async setup(){
    let empty = ref(false)
    let pageNum = ref(1)
    let currentPage = ref(1)
    let blogs = reactive([])
    let response = await getBlogPageNum("", "all")

    pageNum = response.data.pageNum
    response = await getLatestBlogs(1, "all", "")
    if(response.status !== 0){
      empty = true
    }else{
      blogs = response.data
      // for(let i = 0; i < blogs.length; i++){
      //   blogs[i].show = false
      // }
    }

    return {
      empty: empty,
      pageNum: pageNum,
      currentPage: currentPage,
      blogs: blogs
    }
  },
  methods: {
    async pass(item){
      const response = await PassBlogs(item)
      if (response.status !== 0){
        alert(response.msg)
      }else{
        alert(response.msg)
        window.location.reload()
      }
    },
    async remove(item){
      const response = await deleteBlog(item)
      if (response.status !== 0){
        alert(response.msg)
      }else{
        alert(response.msg)
        window.location.reload()
      }
    },
    checkBlog(item){
      window.location = '/blogs/' + item.userId + '/article/details/' + item.blogId
    },
    async getPageBlogs(val){
      const response = await getLatestBlogs(val, "all", "")
      this.blogs = response.data
      // this.refreshBlog()
      this.$forceUpdate()
    },
    handleSizeChange(){

    },
    // refreshBlog(){
    //   for(let i = 0; i < this.blogs.length; i++){
    //     this.blogs[i].show = false
    //   }
    // }
  }
})
</script>

<style scoped>
div#block{
  display: flex;
  justify-content: space-around;
}

div#arrangement{
  flex: 70%;
  position: relative;
  left: 10px
}

div#operation{
  flex: 30%;
  min-width: 280px;
  text-align: right;
  display: flex;
  flex-direction: column-reverse;
}

ul{
  list-style: none;
  text-align: left;
  flex: 80%;
}

p.blog-title{
  font-size: large;
}

div#markdown{
  text-align: left;
  min-height: 255px;
  border-radius: 5px;
  border-color: #99a9bf;
  border-style: solid;
  padding: 3px;
  margin: 5px;
}
</style>
