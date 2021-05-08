<template>
  <div class="main">
    <h2>部屋予約</h2>
      <div>
        部屋
        <Select
          :value.sync="work_room"
          :kv_list="master_list.room_list"
        ></Select>
        期間
        <Select
          :value.sync="reserve_period"
          :kv_list="master_list.room_list"
        ></Select>
        <div>
          カレンダー
        </div>
      </div>    
  </div>
</template>
<script>
import Select from '@/components/Forms/Select.vue';
import Sugar from 'sugar';

export default {
  name: 'Reserve',
  components:{
    Select
  },
  data() {
    return {
      "work_room":"1",
      "reserve_period": "",
      "work_state":"1"
    };
  },
  methods:{

  },
  computed:{
    master_list:{
      get() {
        return this.$store.getters["master_list/get_master_list"];
      }
    }
  },
  mounted() {
  },
  created() {
    this.reserve_period= new Sugar.Date().format('{yyyy}{MM}{dd}').raw;
  }
}
</script>
