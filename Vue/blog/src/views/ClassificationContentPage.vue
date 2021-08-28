<template>
  <div v-if="isLatest" class="blogs-div">
      <div class="tag-div">
        <span class="blogs-tag"> 最新发表 </span>
      </div>
      <ul>
        <li v-for="latestBlog in latestBlogs" class="blog-item">
          <div class="blog-content-div">
            <span class="title"> <a :href="'/blogs/'+latestBlog.userId+'/article/details/'+latestBlog.blogId">{{latestBlog.title}}</a></span>
            <div class="infos">
              <div>
                <span class="author"> {{latestBlog.author}}</span>
                <span class="digest"> {{latestBlog.digest}}</span>
              </div>
              <span class="time"> {{ latestBlog.updateAt.split("T")[0] }}</span>
            </div>
          </div>
        </li>
      </ul>
  </div>

  <div v-if="!isLatest" class="blogs-div">
    <ul>
      <li v-for="tagBlog in tagBlogs" class="blog-item">
        <div class="tag-div">
          <span class="blogs-tag">{{ fields }} > {{ classes }} > {{ tagBlog.tag }} </span> <span>共 {{tagBlog.num}} 篇博文</span>
        </div>
        <ul>
          <li v-for="blog in tagBlog.blogs" class="blog-item">
            <div class="blog-content-div">
              <span class="title"> <a :href="'/blogs/'+blog.userId+'/article/details/'+blog.blogId">{{blog.title}}</a></span>
              <div class="infos">
                <div>
                  <span class="author"> {{blog.author}}</span>
                  <span class="digest"> {{blog.digest}}</span>
                </div>
              </div>
            </div>
          </li>

        </ul>

      </li>

    </ul>
  </div>

</template>

<script>
import {defineComponent, reactive} from "vue";
import {useRoute} from "vue-router";
import {getTagBlogs} from "../api/classification";
import {getLatestBlogs} from "../api/editBlog";

export default defineComponent({
  name: "ClassificationContentPage",
  async setup(){
    const $route = useRoute()
    const classes = $route.params.classes
    const fields = $route.params.field
    let latestBlogs = reactive([])
    let tagBlogs = reactive([])
    const isLatest = classes === "latest"
    if(!isLatest){
      const response = await getTagBlogs("all", "", fields, classes)
      if(response.status !== 0){
        alert(response.msg)
      }else{
        console.log(response.data)
        tagBlogs = reactive(response.data)
      }
    }else{
      const response = await getLatestBlogs(1, "all", "")
      latestBlogs = response.data.slice(0, 15)
    }

    return {
      tagBlogs: tagBlogs,
      isLatest: isLatest,
      latestBlogs: latestBlogs,
      classes: classes,
      fields: fields
    }
  }
})
</script>

<style scoped>


ul{
  list-style: none;
  display: flex;
  flex-direction: column;
  margin: 0;
  padding: 0;
  border-style: solid;
  border-radius: 15px;
  border-color: #99a9bf;
}

div.tag-div{
  text-align: left;
}

div.blog-content-div{
  display: flex;
  text-align: left;
  flex-direction: column;
  padding: 18px 24px 13px 24px;
}

li.blog-item{
  font-family: PingFang SC,Hiragino Sans GB,Arial,Microsoft YaHei,Verdana,Roboto,Noto,Helvetica Neue,sans-serif !important;
  background-color: white;
  border-style: solid;
  border-radius: 10px;
  border-color: #99a9bf;
  border-bottom-color: #3485ff;
}

span.title{
  padding: 3px;
  font-weight: bold;
  font-size: 24px;
}

div.infos{
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

span.author{
  color: #3d3d3d;
  font-size: 16px;
}

span.digest{
  padding-left: 15px;
  font-size: 14px;
  color: #8a8a8a;
}

span.blogs-tag{
  font-size: 20px;
  padding-right: 10px;
  color: #f657ff;
}

</style>
