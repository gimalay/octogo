package com.example.todolist.ui

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import androidx.compose.state
import androidx.compose.unaryPlus
import androidx.ui.core.setContent
import com.example.todolist.model.HomeModel

class MainActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            val homeModel = +state { HomeModel() }
            App(homeModel.value)
        }
    }
}
