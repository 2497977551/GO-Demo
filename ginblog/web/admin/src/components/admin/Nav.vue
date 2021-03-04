<template>
  <a-layout-sider breakpoint="lg" collapsed-width="0" >
    <div class="logo">Blog</div>
    <a-menu theme="dark" mode="inline"  @click="goPage">
      <a-menu-item key="index">
        <a-icon type="dashboard" theme="filled" />
        <span class="nav-text">仪表盘</span>
      </a-menu-item>
      <a-sub-menu>
        <span slot="title"
          ><a-icon type="book" theme="filled" /><span>文章管理</span></span
        >
        <a-menu-item key="addart">
          <a-icon type="edit" theme="filled" />
          <span>添加文章</span>
        </a-menu-item>
        <a-menu-item key="artlist">
          <a-icon type="switcher" theme="filled" />
          <span>文章列表</span>
        </a-menu-item>
      </a-sub-menu>
      <a-menu-item key="catelist">
        <a-icon type="database" theme="filled" />
        <span class="nav-text">分类管理</span>
      </a-menu-item>
      <a-menu-item key="userlist">
        <a-icon type="user" />
        <span class="nav-text">用户列表</span>
      </a-menu-item>
    </a-menu>
    <div id="headerBtn" @click="loginout" style="width:150px">
      <a-button type="danger" block>退出</a-button>
    </div>
  </a-layout-sider>
</template>

<script>
export default {

  methods: {
    loginout() {
      window.sessionStorage.clear('Token')
      this.$router.push('/login')
    },
    goPage(item) {
      this.$router.push('/admin/'+item.key)
      
    }
  }
}
</script>
<style scoped>
.logo {
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
  top: 40%;
  left: 10%;
}
</style>
