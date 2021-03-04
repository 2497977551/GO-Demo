<template>
  <a-layout id="container">
    <Nav></Nav>
    <a-layout>
      <Content></Content>
      <Footer></Footer>
    </a-layout>
  </a-layout>
</template>
<script>
import Nav from '../components/admin/Nav'
import Footer from '../components/admin/Footer'
import Content from '../components/admin/Content'
export default {
  components: { Nav, Footer, Content },
  
}
</script>

<style>
#container {
  height: 100%;
}

</style>
