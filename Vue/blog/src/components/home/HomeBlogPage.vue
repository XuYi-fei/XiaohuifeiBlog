<template>
  <ul id="blog-div" style="overflow: auto; max-height: 1440px">
    <li class="blog-li" v-for="item in blogs">
      <div class="blog-title" style="text-align: center"><a class="title-a" :href="'/blogs/'+ item.userId + '/article/details/' +item.blogId"> >>> {{item.title}} </a></div>
      <el-divider style="margin: 4px"></el-divider>
      <div class="blog-content">
        <div class="blog-text-div">
          <div class="digest"> &nbsp;&nbsp;&nbsp;&nbsp;{{ item.digest }} </div>
          <br>
          <div class="info-detail">
            <div class="blog-info">
              <span class="kind"> {{item.field}} </span>
              >>
              <span class="kind"> {{item.classification}}</span>
              >>
              <span class="kind"> {{item.tag}}</span>
            </div>
            <div class="blog-info">
                  <span>
                    <el-icon :size="15" color="orange">
                      <clock />
                    </el-icon>
                    {{ item.updateAt.split("+")[0].split(".")[0].replace("T", " ") }}
                  </span>
            </div>
            <div class="blog-info">
                    <span>
                      <el-icon :size="15" color="orange">
                        <user-filled />
                      </el-icon>
                      {{ item.author}}
                    </span>
            </div>
          </div>
        </div>
      </div>
    </li>
  </ul>
  <el-pagination
      @size-change="handleSizeChange"
      @current-change="getPageBlogs"
      v-model:currentPage="currentPage"
      layout="prev, pager, next, jumper"
      :page-count="pageNum"
      :key="pageNum"
      style="position: relative;z-index: 30000"
  >
  </el-pagination>
</template>

<script>
import {defineComponent, reactive, ref} from "vue";
import { Edit, UserFilled, Clock } from '@element-plus/icons'
import {getBlogPageNum, getLatestBlogs, getPageNum} from "../../api/editBlog";
export default defineComponent({
  name: "HomeBlogPage",
  components: {
    Edit,
    UserFilled,
    Clock
  },
  async setup(){
    let response = await getBlogPageNum("", "all")
    let currentPage = ref(1)
    let pageNum = ref(1)
    if( response.status !== 0){
      alert("服务器错误,请重试")
      return
    }
    pageNum = response.data.pageNum
    response = await getLatestBlogs(1, "all", "")
    const blogs = reactive(response.data)
    return{
      pageNum: pageNum,
      blogs: blogs,
      currentPage: currentPage
    }
  },
  methods:{
    handleSizeChange(){

    },
    async getPageBlogs(val){
      const response = await getLatestBlogs(val, "all", "")
      this.blogs = response.data
      this.$forceUpdate()
    }
  }
})
</script>

<style scoped>
/* 博客文章 */
a.title-a{
  color: #a1fff5;
  position: relative;
  z-index: 30000;
}

ul#blog-div{
  display: flex;
  list-style: none;
  flex-direction: column;
}

li.blog-li{
  margin: 10px;
  position: relative;
  padding: 5px;
  z-index: 30000;
}

div.blog-title{
  display: flex;
  justify-content: left;
  font-size: x-large;
  font-weight: bold;
  color: #89f7ff;
  padding: 2px 3px 2px 0;
}

div.blog-content{
  display: flex;
}

div.blog-img-div{
  flex: 30%;
  max-width: 200px;
  display: flex;
}

div.blog-text-div{
  margin: 3px;
  flex: 70%;
  display: flex;
  flex-direction: column;
  border-bottom-style: solid;
  border-bottom-color: #42b983;
  justify-content: space-between;
}

div.digest{
  padding: 5px;
  font-family: "Comic Sans MS", cursive, sans-serif;
  align-self: flex-start;
  text-align: left;
  font-size: 20px;
  color: white;
}

img.img{
  display: inline-flex;
}
span.kind{
  color: #ff6b18;
}
div.blog-info{
  padding: 1px 3px 1px 3px;
  text-align: left;
  color: white;
}

</style>
