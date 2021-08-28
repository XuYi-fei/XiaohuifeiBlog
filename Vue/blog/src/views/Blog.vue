<template>
  <vue-particles
      color="#fff"
      :particleOpacity="0.7"
      :particlesNumber="60"
      shapeType="circle"
      :particleSize="4"
      linesColor="#fff"
      :linesWidth="2"
      :lineLinked="true"
      :lineOpacity="0.4"
      :linesDistance="150"
      :moveSpeed="2"
      :hoverEffect="true"
      hoverMode="grab"
      :clickEffect="true"
      clickMode="push"
      class="lizi"
      style="height: 1200px;min-width: 1440px; background: rgba(20 ,20 ,20, 0); width: 100vw; position: absolute; top: 0; z-index: 20000"
  >
  </vue-particles>
  <div style="position: absolute; top: 0; min-width: 1440px; width: 100vw; background-color: black">
    <Header> </Header>
    <div id="main">
      <div style="flex: 10%; max-width: 10%">

      </div>
      <div style="flex: 60%; max-width: 60%;">
        <div id="article-info">
          <div class="nav-title">
            <span> >>> {{ body.title }} </span>
          </div>

          <div id="article-info-content">
          <span class="info-detail">
            <span class="info-detail-icon">
              <el-icon :size="20" color="orange">
                <clock />
              </el-icon>
            </span>
            <span> 创建时间： {{ body.created_time }}</span>
          </span>

            <span class="info-detail">
            <span class="info-detail-icon">
              <el-icon :size="20" color="orange">
                <user-filled />
              </el-icon>
            </span>
            <span class="info-detail-tag">作者： {{ body.author }}</span>
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
        <div id="markdown" style="position: relative; z-index: 30000">
          <v-md-editor v-model="body.content" mode="preview" height="800px" right-toolbar="toc sync-scroll" ></v-md-editor>
        </div>
      </div>
      <div style="flex: 30%; max-width: 30%">
      </div>
    </div>
  </div>
</template>

<script>
import Header from "../components/Header";
import {Clock, UserFilled} from '@element-plus/icons'
import {getBlogById} from "../api/editBlog";
import Markdown from "./Markdown";
export default {
  name: "Blog",
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
        created_time: '',
      },
    }
  },
  async created() {
    try {
      const response = await getBlogById(this.$route.params.blogId)
      if (response.status === 0){
        this.body = response.data
        this.body.created_time = this.body.createdTime.replace('T', ' ').split('.')[0].split('+')[0]
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
  min-width: 1440px;
}

div#article-info{
  display: flex;
  color: white;
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
  text-align: left;
  border-radius: 5px;
  border-color: #99a9bf;
  border-style: solid;
  padding: 3px;
}
</style>
