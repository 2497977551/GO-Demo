<template>
  <a-layout id="container">
    <a-layout-sider breakpoint="lg" collapsed-width="0">
      <div class="logo">Hello</div>
      <Nav></Nav>
      <div id="headerBtn" @click="loginout" style="width:150px">
        <a-button type="danger" block>退出</a-button>
      </div>
    </a-layout-sider>
    <a-layout>
      <a-layout-content :style="{ margin: '12px ' }">
        <div
          :style="{ padding: '24px', background: '#fff', minHeight: '100%' }"
        >
          content
        </div>
      </a-layout-content>
      <a-layout-footer><Footer></Footer></a-layout-footer>
    </a-layout>
  </a-layout>
</template>
<script>
import Nav from '../components/admin/Nav'
import Footer from '../components/admin/Footer'
export default {
  components: { Nav, Footer },
  methods: {
    loginout() {
      window.sessionStorage.clear('Token')
      this.$router.push('/login')
    }
  }
}
</script>

<style>
#container {
  height: 100%;
}
#container .logo {
  height: 32px;
  background: rgba(255, 255, 255, 0.2);
  margin: 16px;
  font-size: 20px;
  background-color: aliceblue;
  display: flex;
  justify-content: center;
}
#headerBtn {
  text-align: center;
  position: relative;
  top: 50%;
  left: 10%;
}
</style>
