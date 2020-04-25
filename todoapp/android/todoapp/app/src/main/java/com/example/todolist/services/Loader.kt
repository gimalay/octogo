package com.example.todolist.services

import com.example.todolist.model.ViewModelOuterClass.ViewModel
import com.example.todolist.model.ViewModelOuterClass.Location

object Loader {
    object Home {
        fun get(): ViewModel.Home {
            val payload = Location.Home.newBuilder().build()
            val data = GoBinding.read(payload)
            return ViewModel.Home.parseFrom(data)
        }
    }

    object Project {
        // ...
    }

}
