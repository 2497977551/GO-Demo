<template>
    <div>
        <h3>用户列表</h3>
        <a-card>
            <a-row :gutter="15">
              <a-col :span="7">
                  <a-input-search placeholder="input search text" enter-button @search="onSearch" />
              </a-col>
              <a-col :span="3">
                   <a-button type="primary">新增用户</a-button>
              </a-col>
            </a-row>
            
            <a-table dataSource="userlist"></a-table>
            
        </a-card>
    </div>
</template>
<script>
const column = []
export default {
    
    data(){
        return {
            userlist :[],
            column
        }
    },
    methods:{
        getUserList(){
            this.$http.get('QueryAllUser',{PageNumber:{},PageSize:{}})
        }
    }
}
</script>