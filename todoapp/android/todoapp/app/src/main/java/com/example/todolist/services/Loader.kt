package com.example.todolist.services

import com.example.todolist.model.ViewModelOuterClass.ViewModel
import com.example.todolist.model.ViewModelOuterClass.Location

class Loader(private val binding: GoBinding) {

    fun getHome(): ViewModel.Home {
        val payload = Location.Home.newBuilder().build()
        val data = binding.read(payload)
        return ViewModel.Home.parseFrom(data)
    }

}
