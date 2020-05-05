package com.example.todolist.di

import com.example.todolist.services.GoBinding
import com.example.todolist.services.Loader
import com.example.todolist.services.Executor

/***
 * An alternative to dependency injection
 * https://developer.android.com/training/dependency-injection#di-alternatives
 */
object ServiceLocator {
    val goBinding: GoBinding by lazy {
        GoBinding()
    }
    val loader: Loader by lazy {
        Loader()
    }
    val executor: Executor by lazy {
        Executor()
    }
}
