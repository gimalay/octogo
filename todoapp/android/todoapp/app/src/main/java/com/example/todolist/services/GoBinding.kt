package com.example.todolist.services

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

    fun read(payload: Message): ByteArray {
        val type = LocationType.valueOf(payload.javaClass.simpleName)
        val location = Location
            .newBuilder()
            .setType(type)
            .setPayload(payload.toByteString())
            .build()

        val data = binding.viewModel(location.toByteArray())

        return data
    }

    fun execute(payload: Message) {
        val ts = com.google.protobuf.Timestamp
            .newBuilder()
            .setSeconds(System.currentTimeMillis() / 1000)
            .build()

        val type = CommandType.valueOf(payload.javaClass.simpleName)
        val command = Command
            .newBuilder()
            .setType(type)
            .setPayload(payload.toByteString())
            .setTimestamp(ts)
            .build()

        binding.execute(command.toByteArray())
    }
}
