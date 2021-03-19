<template>
  <div class="container">
    <div class="loginBox">
      <a-form-model
        ref="loginForm"
        :rules="rules"
        :model="fromData"
        class="loginFrom"
      >
        <a-form-model-item prop="UserName">
          <a-input placeholder="Username" v-model="fromData.UserName">
            <a-icon
              slot="prefix"
              type="user"
              style="color: rgba(0, 0, 0, 0.25)"
            />
          </a-input>
        </a-form-model-item>

        <a-form-model-item prop="PassWord">
          <a-input
            type="password"
            placeholder="Password"
            v-model="fromData.PassWord"
            v-on:keyup.enter="login"
          >
            <a-icon
              slot="prefix"
              type="lock"
              style="color: rgba(0, 0, 0, 0.25)"
            />
          </a-input>
        </a-form-model-item>

        <a-form-model-item class="loginBtn">
          <div style="width: 360px">
            <a-button type="primary" block @click="login"> 登录 </a-button>
          </div>
        </a-form-model-item>
      </a-form-model>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      fromData: {
        UserName: '',
        PassWord: '',
      },
      rules: {
        UserName: [
          {
            required: true,
            message: '用户名不能为空',
            trigger: 'blur',
          },
          {
            min: 4,
            max: 10,
            message: '用户名不可少于4位或大于10位',
            trigger: 'blur',
          },
        ],
        PassWord: [
          {
            required: true,
            message: '密码不能为空',
            trigger: 'blur',
          },
          {
            min: 8,
            max: 16,
            message: '密码不可少于8位或大于16位',
            trigger: 'blur',
          },
        ],
      },
    }
  },
  methods: {
    login() {
      this.$refs.loginForm.validate(async (valid) => {
        if (!valid) {
          return this.$message.error('请按照提示输入用户名与密码！')
        }
        const res = await this.$http.post('Login', this.fromData)
        if (res.data.Status !== 200) {
          return this.$message.warning(res.data.MessAge)
        }
        window.localStorage.setItem('Token', res.data.Token)
        console.log(res.data.Token)
        this.$router.push('admin/index')
        return this.$message.success('登录成功')
      })
    },
  },
}
</script>

<style scoped>
.container {
  height: 100%;
  background-image: url('../assets/bier.jpg');
  background-size: 100% 100%;
}
.loginBox {
  height: 300px;
  width: 400px;
  background-color: seashell;
  position: absolute;
  top: 50%;
  left: 70%;
  box-shadow: papayawhip 0px 0px 10px;
  transform: translate(-50%, -50%);
  border-radius: 4px;
  opacity: 0.75;
}
.loginFrom {
  width: 100%;
  position: absolute;
  padding: 0 20px;
  bottom: 10%;
  box-sizing: border-box;
}
.loginBtn {
  display: flex;
  justify-content: center;
}
</style>
