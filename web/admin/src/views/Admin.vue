<template>
  <a-layout id="container">
    <Nav></Nav>
    <a-layout>
      <a-layout-content>
        <router-view></router-view>
      </a-layout-content>
      <a-layout-footer>
        <Footer></Footer>
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>
<script>
import Nav from '../components/admin/Nav'
import Footer from '../components/admin/Footer'

export default {
  components: { Nav, Footer }
}
</script>

<style>
#container {
  height: 100%;
}
</style>
