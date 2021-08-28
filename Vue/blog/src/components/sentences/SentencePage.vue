<template>
  <div id="contents">
    <div id="contents-div-show">
      <ul id="sentences">
        <li v-for="item in sentences" :key="item">
          <div>
            <span class="content-show sen" style="white-space: pre-line">{{ item.content }}</span>
            <span class="content-show author">{{ item.author === ''? '': '—— ' + item.author }}</span>
          </div>
          <div>
              <span class="user-name">
                <user-filled style="width: 25px; position: relative; top: 5px"/>
                {{ item.userName }}
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                <timer style="width: 25px; position: relative; top: 5px"/>
                {{ item.createdTime.split("+")[0].split(".")[0].replace("T"," ")}}
              </span>
          </div>
          <el-divider style="margin: 10px"></el-divider>
        </li>
      </ul>
    </div>

    <div id="contents-nav-show">
      <el-button round @click="dialogFormVisible = true">挥一笔属于自己的墨</el-button>

      <el-dialog title="属于你的文字" v-model="dialogFormVisible">
        <el-form :model="form" :rules="rules" ref="form">
          <el-form-item label="正文" :label-width="formLabelWidth" prop="content">
            <el-input type="textarea" v-model="form.content"  :resize="formResize"></el-input>
          </el-form-item>
          <el-form-item label="原句作者" :label-width="formLabelWidth" prop="author">
            <el-input type="textarea" v-model="form.author" :resize="formResize"></el-input>
          </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
              <el-button @click="dialogFormVisible = false">取 消</el-button>
              <el-button type="primary" @click="submitForm('form')">确 定</el-button>
            </span>
        </template>
      </el-dialog>
    </div>
  </div>
  <div class="block">
    <el-pagination
        @size-change="handleSizeChange"
        @current-change="getPageSentences"
        v-model:currentPage="currentPage"
        layout="prev, pager, next, jumper"
        :page-count="pageNum"
        >
    </el-pagination>
  </div>
</template>

<script>
import {ref, defineComponent, defineAsyncComponent} from "vue";
import {getAllSentences, getSentencesPageNum, submitSentence} from "../../api/sentences";
import {Timer, UserFilled} from "@element-plus/icons";
import {loginCheck} from "../../api/login";

export default defineComponent({
  name: "SentencePage",
  components: {
    UserFilled,
    Timer
  },
  async setup(){
    const pageResponse = await getSentencesPageNum("all", "")
    const pageNum = pageResponse.data.pageNum
    const currentPage = ref(1)
    const response = await getAllSentences(1, "all", "")
    let sentences = []
    if (response.status === 0){
      sentences = response.data
    }else{
      alert("暂无数据")
    }
    return {
      pageNum: pageNum,
      currentPage: currentPage,
      sentences: sentences
    }
  },
  methods: {
    async uploadSentence(){
      let response = await loginCheck()
      if (response.status !== 0){
        alert("请先登录")
        window.location.reload()
      }else{
        this.body.userId = response.data.userId
        this.body.content = this.form.content
        this.body.author = this.form.author
        response = await submitSentence(this.body)
        if (response.status !== 0 ){
          alert("服务器错误，请重试")
        }else{
          alert(response.msg)
          window.location.reload()
        }
      }
    },
    async getPageSentences(val){
      const response = await getAllSentences(val, "all", "")
      this.sentences = response.data
      this.$forceUpdate()
    },
    handleSizeChange: (val) => {
      // console.log(`每页 ${val} 条`);
    },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.uploadSentence()
        } else {
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    }
  },
  data(){
    return{
      rules: {
        content: [
          { required: true, message: '分享下你的珍藏的句子叭！', trigger: 'blur' }
        ]
      },
      dialogFormVisible: false,
      form: {
        content: '',
        author: ''
      },
      body: {
        author: '',
        content: '',
        userId: ''
      },
      formLabelWidth: '120px',
      formResize: 'none'
    }
  }
})
</script>

<style scoped>
.el-carousel__item div {
  color: #475669;
  font-size: 18px;
  opacity: 0.75;
  line-height: 45px;
  margin: 0;
}

.el-carousel__item:nth-child(2n) {
  background-color: #fce0ff;
}

.el-carousel__item:nth-child(2n+1) {
  background-color: #e4ffe0;
}

div#carousel{
  margin-top: -20px;
}

div#contents{
  min-width: 1440px;
  display: flex;
  margin: 10px;
}

div#contents-div-show{
  flex: 75%;
}

div#contents-nav-show{
  margin-left: 30px;
  margin-right: 50px;
  min-width: 160px;
  flex: 20%;
  display: flex;
  flex-direction: column;
  position: relative;
  z-index: 30000;
}

ul{
  text-align: left;
  list-style-type: none;
}

span.content-show{
  display: block;
  color: white;
}

span.sen{
  padding: 0 10px 0 10px;
  text-align: left;
  color: white;
  font-size: 20px;
}

span.author{
  text-align: right;
  font-size: large;
  padding-bottom: 10px;
}

span.user-name{
  color: orange;
  padding: 0 10px 0 10px;
  font-size: 18px;
}
</style>
