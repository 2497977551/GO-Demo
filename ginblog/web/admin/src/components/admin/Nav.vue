<template>
    <a-menu theme="dark" mode="inline" :default-selected-keys="['1']">
        <a-menu-item key="1">
          <a-icon type="dashboard" theme="filled" />
          <span class="nav-text">1</span>
        </a-menu-item>
        <a-sub-menu key="sub1">
        <span slot="title"><a-icon type="book" theme="filled" /><span>2</span></span>
        <a-menu-item key="5">
            <a-icon type="edit" theme="filled" />
         <span>1</span>
        </a-menu-item>
        <a-menu-item key="6">
            <a-icon type="switcher" theme="filled" />
          <span>2</span>
        </a-menu-item>
      
      </a-sub-menu>
        <a-menu-item key="3">
          <a-icon type="database" theme="filled" />
          <span class="nav-text">3</span>
        </a-menu-item>
        <a-menu-item key="4">
          <a-icon type="user" />
          <span class="nav-text">4</span>
        </a-menu-item>
        
      </a-menu>
      
</template>