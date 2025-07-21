package com.ticketfaster.orders.service;


import com.ticketfaster.orders.model.Order;
import com.ticketfaster.orders.repository.OrderRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class OrderService {

    @Autowired
    private OrderRepository orderRepository;

    public List<Order> getAllOrders() {
        return orderRepository.findAll();
    }

    public Optional<Order> getOrderById(Long id) {
        return orderRepository.findById(id);
    }

    public Order createOrder(Order order) {
        // You would add business logic here, e.g., setting initial status
        order.setStatus("CREATED");
        return orderRepository.save(order);
    }
}