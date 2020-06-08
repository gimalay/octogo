package com.example.todolist.service

import android.os.Environment
import binding.Binding
import binding.Binding_
import com.example.todolist.model.CommandOuterClass.Command
import com.example.todolist.model.CommandOuterClass.CommandType
import com.example.todolist.model.ViewModelOuterClass.Location
import com.example.todolist.model.ViewModelOuterClass.LocationType
import com.google.protobuf.Message
import java.io.File

class GoBinding {
    private var binding: Binding_

    init {
        val path: String = Environment
            .getExternalStorageDirectory()
            .getAbsolutePath()

        val dbDirPath = "$path/Android/data/com.example.todolist"
        val dbDir = File(dbDirPath)
        dbDir.mkdirs()

        binding = Binding.new_("$dbDirPath/boltdb")
    }

    fun read(payload: Message): ByteArray? {
        return binding.viewModel(payload.toByteArray())
    }

    fun execute(payload: Message) {
        binding.execute(payload.toByteArray())
    }

    fun close() {
        binding.close()
    }
}
