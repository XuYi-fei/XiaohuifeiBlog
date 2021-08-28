<template>
  <el-empty :image-size="300" v-show="empty" style="height: 720px" description="暂无数据"></el-empty>
  <ul id="articles">
    <li v-for="item in blogInfo" :key="item">
      <div class="blog-content">
        <div class="title"> {{item.title}} </div>
        <div class="digest"> {{ item.digest }} </div>
        <div class="detail-info">
          <ul class="infos">
            <li class="info date"> <span > 发布日期: </span> <span>{{ item.updateAt.replace("T", " ").split("+")[0].split(".")[0]}} </span></li>
          </ul>
          <ul class="operations">
            <li class="info"><a :href="'/personal/'+ user.userId + '/editblog/update/' + item.blogId"> 编辑 </a></li>
            <li class="info"><a @click.prevent="checkBlog(item.blogId)"> 浏览 </a></li>
            <li class="info"><a @click.prevent="removeBlog(item)"> 删除 </a></li>
          </ul>
        </div>
      </div>
      <el-divider style="margin: 4px"></el-divider>
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
import {deleteBlog, getBlogPageNum, getLatestBlogs} from "../../api/editBlog";
import {loginCheck} from "../../api/login";

export default defineComponent({
  name: "EditBlogPage",
  async setup(){
    let response = await loginCheck()
    if (response.status !== 0){
      alert("登录失效，请重新登录")
      window.location = '/home'
    }
    let user = response.data
    let currentPage = ref(1)
    let pageNum = ref(1)
    let blogInfo = reactive([])
    let empty = false
    response = await getBlogPageNum(user.userId, "personal")
    pageNum = response.data.pageNum
    response = await getLatestBlogs(1, "personal", user.userId)
    if (response.status !== 0){
      empty = true
    }else{
      blogInfo = response.data
    }
    return{
      pageNum: pageNum,
      blogInfo: blogInfo,
      currentPage: currentPage,
      user: user,
      empty: empty
    }
  },
  methods:{
    handleSizeChange(){

    },
    async getPageBlogs(val){
      const response = await getLatestBlogs(val, "personal", this.user.userId)
      this.blogInfo = response.data
      this.$forceUpdate()
    },
    checkBlog(BlogId){
      window.location = '/blogs/' + this.user.userId + '/article/details/' + BlogId
    },
    async removeBlog(item){
      const response = await deleteBlog(item)
      if (response.status === 0){
        alert("删除成功!")
        window.location.reload()
      }else{
        alert(response.msg)
      }
    }
  }
})
</script>

<style scoped>
ul#articles{
  display: flex;
  flex-direction: column;
  list-style: none;
}

li{
  display: block;
}

div.blog-content{
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

div.title{
  font-size: large;
  padding-bottom: 30px;
  padding-left: 10px;
  font-weight: bold;
}

div.digest{
  padding-top: 0;
  padding-left: 15px;
  padding-bottom: 30px;
}

div.detail-info{
  display: flex;
  justify-content: space-between;
}

ul.infos{
  list-style: none;
  padding: 4px;
}

li.info{
  display: inline-block;
  padding: 5px;
}

li.date{
  color: grey;
}
</style>
