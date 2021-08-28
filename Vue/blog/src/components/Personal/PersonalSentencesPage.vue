<template>
  <div id="main">
    <el-empty :image-size="300" v-show="empty" style="height: 740px" description="暂无数据"></el-empty>
    <ul id="articles">
      <li v-for="item in sentences" :key="item">
        <div class="blog-content">
          <div class="title" style="white-space: pre-line"> {{item.content}} </div>
          <div class="detail-info">
            <ul class="infos">
              <li class="info date"> <span > 发布日期: </span> <span>{{ item.createdTime.replace("T", " ").split("+")[0].split(".")[0]}} </span></li>
            </ul>
            <ul class="operations">
              <li class="info"><a @click.prevent="removeSentence(item)"> 删除 </a></li>
            </ul>
          </div>
        </div>
        <el-divider style="margin: 4px"></el-divider>
      </li>
    </ul>
    <div>
      <el-pagination
          @size-change="handleSizeChange"
          @current-change="getPageSentences"
          v-model:currentPage="currentPage"
          layout="prev, pager, next, jumper"
          :page-count="pageNum"
          :key="pageNum"
      >
      </el-pagination>
    </div>
  </div>
</template>

<script>
import {defineComponent, reactive, ref} from "vue";
import {loginCheck} from "../../api/login";
import {
  DeleteSentencesPersonal, getAllSentences,
  getSentencesPageNum,
} from "../../api/sentences";

export default defineComponent({
  name: "PersonalSentencesPage",
  async setup(){
    let sentences = reactive([])
    const empty = ref(false)
    let response = await loginCheck()
    if (response.status !== 0){
      alert("登录失效，请重新登录")
      window.location = '/home'
    }
    const user = response.data
    response = await getSentencesPageNum("personal", user.userId)
    if(response.status !== 0){
      alert("未知错误")
    }
    const pageNum = ref(response.data.pageNum)
    const currentPage = ref(1)
    response = await getAllSentences(1, "personal", user.userId)
    if (response.status !== 0){
      this.empty = true
    }else{
      sentences = response.data
    }
    return {
      pageNum: pageNum,
      currentPage: currentPage,
      user: user,
      sentences: sentences,
      empty: empty
    }
  },
  methods:{
    async removeSentence(item){
      const response = await DeleteSentencesPersonal(item)
      if (response.status === 0){
        alert("删除成功!")
        window.location.reload()
        const response = await getSentencesPageNum("personal", this.user.userId)
        this.pageNum = response.data.pageNum
        this.$forceUpdate()
      }else{
        alert(response.msg)
      }
    },
    async getPageSentences(val){
      const response = await getAllSentences(val,"personal", this.user.userId)
      if (response.status !== 0){
        this.empty = true
      }else{
        this.sentences = response.data
        this.$forceUpdate()
      }
    },
    handleSizeChange(){

    }
  }
})
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
