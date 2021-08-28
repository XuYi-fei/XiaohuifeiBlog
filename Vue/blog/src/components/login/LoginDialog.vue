<template>
  <el-button type="text" @click="dialogVisible = true" show-close="true"> 登录 </el-button>

  <div class="el-overlay" style="z-index: 2002" v-show="dialogVisible">
    <div class="el-dialog" aria-modal="true" role="dialog" aria-label="登录" style="margin-top: 15vh;min-width: 500px; width: 770px; height: 350px ; display: flex; flex-direction: column; align-items: center">
      <div class="el-dialog__header">
        <span class="el-dialog__title">登录</span>
        <button aria-label="close" class="el-dialog__headerbtn" type="button">
          <i class="el-dialog__close el-icon el-icon-close" @click="dialogVisible = false"></i>
        </button>
      </div>
      <div class="el-dialog__body" style="width: 50%; min-width: 100px; position: relative; right: 5%">
        <div>
          <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
            <el-form-item label="账号" prop="userId">
              <el-input v-model="ruleForm.userId"></el-input>
            </el-form-item>
            <el-form-item label="密码" prop="pass">
              <el-input type="password" v-model="ruleForm.pass" autocomplete="off"></el-input>
            </el-form-item>
          </el-form>
        </div>
      </div>
      <a href="/forgettingPassword" style="position: relative; bottom: 45px">忘记密码?</a>
      <div class="el-dialog__footer" style="align-self: flex-end">
        <span class="dialog-footer" data-v-1458a6eb="" >
          <button class="el-button el-button--default" type="button" data-v-1458a6eb="" @click="dialogVisible = false">
            <span>取 消</span>
          </button>
          <button class="el-button el-button--default" type="button" data-v-1458a6eb="" @click="resetForm('ruleForm')">
            <span>重 置</span>
          </button>
          <button class="el-button el-button--primary" type="button" data-v-1458a6eb="" @click="submitForm('ruleForm')">
            <span>登 录</span>
          </button>
        </span>
      </div>
    </div>
  </div>

</template>

<script>
import { defineComponent, ref } from 'vue';
import { ElMessageBox } from 'element-plus';
import {postLogin} from "../../api/login";
export default defineComponent({
  name: "LoginDialog",
  setup() {
    const dialogVisible = ref(false);
    const username = ''
    const password = ''
    const userId = ''
    const handleClose = (done) => {
      ElMessageBox
          .confirm('确认关闭？')
          .then((_) => {
            done();
          })
          .catch((_) => {
            // catch
          });
    };
    return {
      dialogVisible,
      handleClose,
      username,
      password,
      userId,
      postLogin
    };
  },
  data () {
    const checkUserId = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('账号不能为空'))
      }
      callback()
    }
    const validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'))
      }
      callback()
    }
    return {
      ruleForm: {
        pass: '',
        userId: ''
      },
      rules: {
        pass: [
          { validator: validatePass, trigger: 'blur' }
        ],
        userId: [
          { validator: checkUserId, trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    submitForm (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.login()
        } else {
          console.log('请正确填写所有字段!!')
          return false
        }
      })
    },
    resetForm (formName) {
      this.$refs[formName].resetFields()
    },
    async login(){
      const response = await postLogin(this.ruleForm)
      if(response.status !== 0){
        alert("用户名或密码错误!!!")
      }
    }
  }
});
</script>

<style scoped>
form{
  float: left;
}


</style>
