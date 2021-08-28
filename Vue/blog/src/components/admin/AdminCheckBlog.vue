<template>
  <div id="main">
    <div style="flex: 10%; max-width: 10%">
    <div style="flex: 60%; max-width: 60%;">
      <div id="article-info">
        <div class="nav-title">
          <span> {{ body.title }} </span>
        </div>

        <div id="article-info-content">
          <span class="info-detail">
            <span class="info-detail-icon">
              <el-icon :size="20" color="orange">
                <clock />
              </el-icon>
            </span>
          </span>

          <span class="info-detail">
            <span class="info-detail-icon">
              <el-icon :size="20" color="orange">
                <user-filled />
              </el-icon>
            </span>
            <span class="info-detail-tag">{{ body.author }}</span>
          </span>

          <span class="info-detail">
            <span class="info-detail-tag">范围: {{ body.field }}</span>
          </span>
          <span class="info-detail">
            <span class="info-detail-tag">类别: {{ body.classification }} </span>
          </span>
          <span class="info-detail">
            <span class="info-detail-tag">标签: {{ body.tag }}</span>
          </span>
        </div>
      </div>

      <el-divider></el-divider>
      <div id="markdown">
        <v-md-editor v-model="body.content" mode="preview" height="800px" right-toolbar="toc sync-scroll"></v-md-editor>
      </div>
    </div>
    <div style="flex: 30%; max-width: 30%">
    </div>
  </div>
  </div>
</template>

<script>
import Header from "../../components/Header";
import {Clock, UserFilled} from '@element-plus/icons'
import Markdown from "../../views/Markdown";
import {getBlogTmp} from "../../api/admin";
export default {
  name: "AdminCheckBlog",
  props: ['userId', 'blogId'],
  components: {
    Header,
    Markdown,
    Clock,
    UserFilled
  },
  data(){
    return{
      body: {
        title: '',
        digest: '',
        content: '',
        author: '',
        field: '',
        classification: '',
        tag: '',
      },
    }
  },
  async created() {
    try {
      const response = await getBlogTmp(this.$route.params.blogId)
      if (response.status === 0){
        this.body = response.data
        console.log(this.body)
      }else{
        alert("博客不存在!")
        window.location = '/home'
      }
    }catch (e) {
      console.log(e)
    }
  },
}
</script>

<style scoped>
code.language-css {
  background-color: #99a9bf;
}

div#main{
  display: flex;
}

div#article-info{
  display: flex;
  flex-direction: column;
  text-align: left;
  font-weight: bold;
  font-size: xx-large;
  font-family: "STKaiti", sans-serif;;
}

div#article-info-content{
  padding: 15px;
  text-align: left;
}

span.info-detail{
  display: block;
  font-weight: normal;
  font-size: medium;
  padding: 2px 0 2px 0 ;
  font-family: "Microsoft YaHei UI", sans-serif;
}

span.info-detail-icon{
  position: relative;
  top: 4px;
}

span.info-detail-tag{
  font-size: large;
  padding-left: 5px;
}

div#markdown{
  z-index: 1000;
  text-align: left;
  height: 800px;
  background-color: #3485ff;
  border-radius: 5px;
  border-color: #99a9bf;
  border-style: solid;
  padding: 3px;
}
</style>
