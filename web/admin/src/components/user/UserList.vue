<template>
  <div>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search placeholder="输入用户名查找" enter-button />
        </a-col>
        <a-col :span="4">
          <a-button type="primary">新增</a-button>
        </a-col>
      </a-row>

      <a-table
        rowKey="Id"
        :columns="columns"
        :pagination="pagination"
        :dataSource="userlist"
        bordered
      >
        <span slot="Role" slot-scope="Role">{{
          Role == true ? '管理员' : '会员'
        }}</span>
        <template slot="action">
          <div>
            <a-button type="primary" style="margin:10px">编辑</a-button>
            <a-button type="danger">删除</a-button>
          </div>
        </template>
      </a-table>
    </a-card>
  </div>
</template>
<script>
const columns = [
  {
    title: 'ID',
    width: '20%',
    dataIndex: 'Id',
    align: 'center',
    key: 'Id'
  },
  {
    title: '用户名',
    width: '20%',
    dataIndex: 'UserName',
    align: 'center',
    key: 'UserName'
  },
  {
    title: '角色',
    width: '10%',
    dataIndex: 'Role',
    align: 'center',
    key: 'admin',
    scopedSlots: { customRender: 'Role' }
  },
  {
    title: '创建时间',
    width: '30%',
    dataIndex: 'CreationTime',
    align: 'center',
    key: 'timer'
  },
  {
    title: '操作',
    width: '20%',
    align: 'center',
    key: 'action',
    scopedSlots: { customRender: 'action' }
  }
]
export default {
  data() {
    return {
      pagination: {
        defaultCurrent: 1,
        defaultPageSize: 20,
        total: 0,
        pageSizeOptions: ['20', '50', '100', '200'],
        showSizeChanger: true,
        showTotal: total => `共${total}条`,
        onChage: (page, pageSize) => {
          this.pagination.defaultCurrent = page
          this.pagination.defaultPageSize = pageSize
        },
        onShowSizeChange: (current, size) => {
          this.pagination.defaultCurrent = current
          this.pagination.defaultPageSize = size
        }
      },
      userlist: [],
      columns
    }
  },
  created() {
    this.getUserList()
  },
  methods: {
    async getUserList() {
      const { data: res } = await this.$http.get('QueryAllUser', {
        params: {
          PageNumber: this.pagination.defaultCurrent,
          PageSize: this.pagination.defaultPageSize
        }
      })

      if (res.Status !== 200) {
        return this.$message.error(res.MessAge)
      } else {
        this.userlist = res.Data.UserInfo

        this.pagination.total = res.Data.Total
      }
    }
  }
}
</script>
