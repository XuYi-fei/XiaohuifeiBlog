<template>
  <el-button type="text" @click="dialogVisible = true" show-close="true"> 注册 </el-button>

  <div class="el-overlay" style="z-index: 2002" v-show="dialogVisible">
    <div class="el-dialog" aria-modal="true" role="dialog" aria-label="注册" style="margin-top: 15vh;min-width: 300px; width: 55%; height: 600px ; display: flex; flex-direction: column; align-items: center">
      <div class="el-dialog__header">
        <span class="el-dialog__title">注册账号</span>
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
            <el-form-item label="用户名" prop="userName">
              <el-input v-model="ruleForm.userName"></el-input>
            </el-form-item>
            <el-form-item label="密保问题" prop="question">
              <el-input v-model="ruleForm.question" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="密保答案" prop="answer">
              <el-input v-model="ruleForm.answer" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="密码" prop="pass">
              <el-input type="password" v-model="ruleForm.pass" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="确认密码" prop="checkPass">
              <el-input type="password" v-model="ruleForm.checkPass" autocomplete="off"></el-input>
            </el-form-item>
          </el-form>
        </div>
      </div>
      <div class="el-dialog__footer" style="align-self: flex-end">
        <span class="dialog-footer" data-v-1458a6eb="" >
          <button class="el-button el-button--default" type="button" data-v-1458a6eb="" @click="dialogVisible = false">
            <span>取 消</span>
          </button>
          <button class="el-button el-button--default" type="button" data-v-1458a6eb="" @click="resetForm('ruleForm')">
            <span>重 置</span>
          </button>
          <button class="el-button el-button--primary" type="button" data-v-1458a6eb="" @click="submitForm('ruleForm')">
            <span>注 册</span>
          </button>
        </span>
      </div>
    </div>
  </div>
</template>

<script>
import { postLogin, postRegister } from '../../api/login'
import { defineComponent, ref } from 'vue';
import { ElMessageBox } from 'element-plus';
export default defineComponent({
  name: "RegisterDialog",
  setup() {
    const dialogVisible = ref(false);
    const username = ''
    const password = ''
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
      password
    };
  },
  data () {
    const checkUserId = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('账号不能为空'))
      }
      callback()
    }
    const checkUserName = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('用户名不能为空'))
      }
      setTimeout(() => {
        callback()
      }, 1000)
    }
    const checkQuestion = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('密保问题用于找回密码，不能为空!'))
      }
      setTimeout(() => {
        callback()
      }, 1000)
    }
    const validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'))
      } else {
        if (this.ruleForm.checkPass !== '') {
          this.$refs.ruleForm.validateField('checkPass')
        }
        callback()
      }
    }
    const validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'))
      } else if (value !== this.ruleForm.pass) {
        callback(new Error('两次输入密码不一致!'))
      } else {
        callback()
      }
    }
    return {
      ruleForm: {
        pass: '',
        checkPass: '',
        userId: '',
        userName: '',
        question: '',
        answer: ''
      },
      rules: {
        pass: [
          { validator: validatePass, trigger: 'blur' }
        ],
        checkPass: [
          { validator: validatePass2, trigger: 'blur' }
        ],
        userId: [
          { validator: checkUserId, trigger: 'blur' }
        ],
        userName: [
          { validator: checkUserName, trigger: 'blur' }
        ],
        question: [
          { validator: checkQuestion, trigger: 'blur' }
        ],
        answer: [
          { validator: checkQuestion, trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    submitForm (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.register()
        } else {
          alert("请正确填写信息")
          return false
        }
      })
    },
    resetForm (formName) {
      this.$refs[formName].resetFields()
    },
    async register(){
      const response = await postRegister(this.ruleForm)
      if (response.status !== 0){
        alert(response.data.msg)
      }else{
        alert("注册成功")
        window.location.reload()
      }
    }
  }
});
</script>

<style scoped>

</style>
