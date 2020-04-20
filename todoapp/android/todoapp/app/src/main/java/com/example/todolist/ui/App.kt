package com.example.todolist.ui

import androidx.compose.Composable
import androidx.compose.onActive
import androidx.compose.state
import androidx.compose.unaryPlus
import androidx.ui.foundation.VerticalScroller
import androidx.ui.layout.Column
import androidx.ui.material.MaterialTheme
import androidx.ui.tooling.preview.Preview
import com.example.todolist.model.CommandOuterClass.Command
import com.example.todolist.model.CommandOuterClass.CommandType
import com.example.todolist.model.ViewModelOuterClass.Location
import com.example.todolist.model.ViewModelOuterClass.LocationType
import com.example.todolist.model.ViewModelOuterClass.ViewModel
import com.example.todolist.services.GoBinding
import com.google.protobuf.ByteString
import java.nio.ByteBuffer
import java.util.*

@Preview
@Composable
fun App() {
    val activitiesModel = +state { ViewModel.Home.getDefaultInstance() }
    +onActive {
        activitiesModel.value = fetch()
    }

    MaterialTheme(
        colors = lightThemeColors,
        typography = themeTypography
    ) {
        VerticalScroller {
            Column {
                activitiesModel.value.projectsList.forEach {
                    Activity(
                        text = it.name,
                        onCopy = {
                            addProject(it.name)
                            activitiesModel.value = fetch()
                        },
                        onRemove = { /*activitiesModel.value.remove(it.id)*/ }
                    )
                }
            }
        }
    }
}

fun newUUIDasByteString(): ByteString? {
    val bb: ByteBuffer = ByteBuffer.wrap(ByteArray(16))
    val uuid = UUID.randomUUID()
    bb.putLong(uuid.mostSignificantBits)
    bb.putLong(uuid.leastSignificantBits)

    return ByteString.copyFrom(bb.array())
}


fun addProject(name: String): Unit {
    val newProject = Command.NewProject
        .newBuilder()
        .setName(name)
        .setProjectID(newUUIDasByteString())
        .build()

    val ts = com.google.protobuf.Timestamp
        .newBuilder()
        .setSeconds(System.currentTimeMillis() / 1000)
        .build()

    val command = Command
        .newBuilder()
        .setType(CommandType.NewProject)
        .setPayload(newProject.toByteString())
        .setTimestamp(ts)
        .build()

    GoBinding.binding.execute(command.toByteArray())
}

fun fetch(): ViewModel.Home {
    val location = Location
        .newBuilder()
        .setType(LocationType.Home)
        .setPayload(Location.Home.newBuilder().build().toByteString())
        .build()

    val data = GoBinding.binding.viewModel(location.toByteArray())

    val home = ViewModel.Home.parseFrom(data)

    return home
}