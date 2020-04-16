package com.example.todolist.model

import androidx.compose.Model
import kotlin.random.Random

@Model
class ActivitiesModel(v: List<Activity>) {

    var items: List<Activity> = v

    fun add(v: String) {
        items = items + Activity(
            id = Random.nextInt(0, 1000),
            name = v
        )
    }

    fun remove(v: Int) {
        items = items.filter { it.id != v }
    }
}

data class Activity(
    val id: Int,
    val name: String
)