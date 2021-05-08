<template >
  <div class="main">
    <h2>出勤報告</h2>
      <div>勤務状態</div>
      <div>
        部屋
        <Select
          :value.sync="work_room"
          :kv_list="master_list.room_list"
        ></Select>
      </div>
      <div>
        <div>報告</div>
        <textarea name="" id="" cols="30" rows="10"></textarea>
      </div>
      <div>
        <button>
          {{button_message}}
        </button>
      </div>
  </div>
</template>
<script>

import Select from '@/components/Forms/Select.vue';

export default {
  name: 'Work',
  components:{
    Select
  },
  data() {
     return {
      "work_room":"1",
      "button_message": "",
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
    this.button_message = (this.work_state === "1") ? "保存します。" : "終了します";
  }
}
</script>
