import Vue from 'vue'
import { Button, FormModel, Input,Icon,message,Layout,Menu,Card,Row,Col,Table,Pagination,ConfigProvider } from 'ant-design-vue'

message.config({
    top: `50px`,
    duration: 3,
    maxCount: 3,
  });
Vue.prototype.$message = message
Vue.use(Button)
Vue.use(FormModel)
Vue.use(Input)
Vue.use(Icon)
Vue.use(Layout)
Vue.use(Menu)
Vue.use(Card)
Vue.use(Row)
Vue.use(Col)
Vue.use(Table)
Vue.use(Pagination)
Vue.use(ConfigProvider)