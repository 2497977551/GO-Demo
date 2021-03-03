<template>
  <div class="container">
    <div class="loginBox">
      <a-form-model :rules="rules" :model="fromData" class="loginFrom">
        <a-form-model-item prop="username">
          <a-input placeholder="Username" v-model="fromData.username">
            <a-icon
              slot="prefix"
              type="user"
              style="color: rgba(0, 0, 0, 0.25)"
            />
          </a-input>
        </a-form-model-item>

        <a-form-model-item prop="password">
          <a-input
            type="password"
            placeholder="Password"
            v-model="fromData.password"
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
            <a-button type="primary" block> 登录 </a-button>
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
        username: '',
        password: '',
      },
      rules: {
        username: [
          {
            required: true,
            message: 'Please input Activity UserName',
            trigger: 'blur',
          },
          {
            min: 4,
            max: 10,
            message: 'Length should be 4 to 10',
            trigger: 'blur',
          },
        ],
        password: [
          {
            required: true,
            message: 'Please input Activity PassWord',
            trigger: 'blur',
          },
          {
            min: 8,
            max: 16,
            message: 'Length should be 8 to 16',
            trigger: 'blur',
          },
        ],
      },
    }
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