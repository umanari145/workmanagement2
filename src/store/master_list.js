
const master_list = {
    namespaced: true,
    state:{
        master_list:{
            "room_list":''
        }
    },
    getters:{
        get_master_list(state) {
            let room_list = {"1":"市川", "2":"福岡"};
            state.master_list.room_list = room_list;
            return state.master_list;
        },
    },
    mutations:{
        set_master_list(state, master_list) {
            state.master_list = master_list;
        },
    },
    actions:{

    }
}

export default master_list
