<template>
  <div id="main">
    <div id="tips1"> <span id="blog-title">博客正文</span></div>
    <el-divider></el-divider>

    <el-form :model="ruleForm" :rules="rules" ref="ruleForm" class="demo-ruleForm" style="width: 80%;padding: 10px">
      <el-form-item label="文章标题" prop="title">
        <el-input v-model="ruleForm.title" ></el-input>
      </el-form-item>
      <el-form-item label="文章摘要" prop="digest">
        <el-input v-model="ruleForm.digest" ></el-input>
      </el-form-item>
      <el-form-item label=">内容方向">
        <el-input v-model="ruleForm.field" ></el-input>
      </el-form-item>
      <el-form-item label=">文章分类">
        <el-input v-model="ruleForm.classification" ></el-input>
      </el-form-item>
      <el-form-item label=">文章  tag">
        <el-input v-model="ruleForm.tag" ></el-input>
      </el-form-item>
    </el-form>

    <div style="text-align: left">
      <v-md-editor v-model="ruleForm.content" height="550px" :include-level="includeLevel" :disabled-menus="[]"  @upload-image="handleUploadImage"></v-md-editor>
    </div>


    <div id="buttons">
      <div id="button-content">
        <el-row>
          <el-button type="success" plain @click="submitForm('ruleForm')"> 点击发表 </el-button>
        </el-row>
      </div>
    </div>


  </div>

</template>

<script>
import VmMarkdown from "vm-markdown"
import "highlight.js/styles/github.css"
import hljs from 'highlight.js'
import {loginCheck} from "../../api/login";
import {submitBlog} from "../../api/editBlog";
import {uploadFile, uploadFiles} from "../../api/files";
export default {
  name: "CreateBlog",
  components: {
    VmMarkdown
  },
  data(){
    return{
      includeLevel: [1, 2, 3],
      ruleForm: {
        title: '',
        digest: '',
        content: '',
        tag: '',
        classification: '',
        userId: ''
      },
      rules: {
        title: [
          { required: true, message: '请输入文章标题', trigger: 'blur' },
          { min: 1, max: 100, message: '文章标题长度过长', trigger: 'blur' }
        ],
        digest: [
          { required: true, message: '请输入文章摘要', trigger: 'blur' }
        ],
      }
    };
  },
  methods: {
    // your method to upolad the file to server
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          // console.log(this.text)
          // console.log(this.ruleForm)
          submitBlog(this.ruleForm)
        } else {
          console.log('error submit!!');
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    async handleUploadImage(event, insertImage, files) {
      // 拿到 files 之后上传到文件服务器，然后向编辑框中插入对应的内容

      for(let i = 0; i < files.length; i ++){
        let params = new FormData()
        params.append("uploadFiles", files[i])
        const response = await uploadFile(params)
            .then(res => {
              return res
            });
        // 此处只做示例
        // 后台返回的状态码，自行根据后台的要求进行修改
        if (response.status === 0) {
          insertImage({
            url: response.data.filePath,
            desc: '在此添加描述',
            // width: 'auto',
            // height: 'auto',
          });
        }
      }

    },
  },
  async created() {
    const response = await loginCheck()
    try{
      if(response.status === 0)
        this.ruleForm.userId = response.data.userId
    }catch (e) {
      alert("登录失效,请重新登录")
      window.location = "/home"
    }
  }
}
</script>

<style scoped>
/* 全局设置 */
div#main{
  display: flex;
  flex-direction: column;
  background-color: white;

}

a{
  text-decoration: none;
}

header{
  padding-top: 8px;
  background-color: white;
  position: relative;
}

/* 组件设置 */
div#BreadCrumb{
  display: flex;
  justify-content: center;
  flex-direction: column;
  padding: 10px;
  flex: 25%;
}


/* 首页导航栏 */
div#header-parent{
  display: flex;
  justify-content: space-around;
}


/* 首页标题部分 */
a#header-title-a{
  display: flex;
  flex-direction: column;
  justify-content: center;
  flex: 50%
}

span#header-title{
  font-family: "Microsoft YaHei UI",serif;
  color: dodgerblue;
  font-size: xx-large;
}

/* 首页链接部分 */
ul#header-ul{
  list-style: none;
  display: inline;
  min-width: 200px;
  white-space: nowrap;
  flex: 25%
}

li{
  display: inline;
  margin: 20px;
}

a.header-a{
  color: dimgrey;
  font-size: large;
}

span#blog-title{
  float:left;
  padding: 0 10px 0 10px
}

/* 按钮 */
div#buttons{
  display: flex;
  flex-direction: row-reverse;
  padding: 4px;
}


</style>
